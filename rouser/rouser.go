package rouser

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"example.com/robloxapi"
)

type Rouser struct {
	Client    *http.Client
	Cookie    string
	TFASecret string
	CSRFToken string

	User *robloxapi.AuthenticatedUser
}

func New(cookie string, tfaSecret string) (*Rouser, error) {
	r := &Rouser{
		Client:    &http.Client{Timeout: 10 * time.Second},
		Cookie:    fmt.Sprintf(".ROBLOSECURITY=_|WARNING:-DO-NOT-SHARE-THIS.--Sharing-this-will-allow-someone-to-log-in-as-you-and-to-steal-your-ROBUX-and-items.|_%s", cookie),
		TFASecret: tfaSecret,
		CSRFToken: "",
	}
	var err error
	r.User, err = r.FetchAuthenticatedUser()
	if err != nil {
		return nil, fmt.Errorf("error fetching authenticated user: %w", err)
	}
	return r, nil
}

func (r *Rouser) GetHeader(solveTFAChallengeRes ...*SolveTFAChallengeRes) http.Header {
	header := http.Header{
		"Content-Type": {"application/json"},
		"Cookie":       {r.Cookie},
		"X-Csrf-Token": {r.CSRFToken},
	}
	if len(solveTFAChallengeRes) > 0 && solveTFAChallengeRes[0] != nil {
		header.Set("Rblx-Challenge-Id", solveTFAChallengeRes[0].RblxChallengeId)
		header.Set("Rblx-Challenge-Metadata", solveTFAChallengeRes[0].RblxChallengeMetadata)
		header.Set("Rblx-Challenge-Type", solveTFAChallengeRes[0].RblxChallengeType)
	}
	return header
}

func (r *Rouser) Request(method string, url string, body []byte) (*robloxapi.Response, error) {
	var solveTFAChallengeRes *SolveTFAChallengeRes
	for {
		req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
		if err != nil {
			return nil, fmt.Errorf("error forming request: %w", err)
		}
		req.Header = r.GetHeader(solveTFAChallengeRes)
		resp, err := robloxapi.Request(req, r.Client)
		if err != nil {
			return nil, fmt.Errorf("error making request: %w", err)
		}
		if resp.APIError != nil {
			if resp.APIError.Message == "Token Validation Failed" {
				r.CSRFToken = resp.Header.Get("X-Csrf-Token")
				continue
			}
			if resp.APIError.Message == "InternalServerError" {
				continue
			}
			if resp.APIError.Message == "Challenge is required to authorize the request" && resp.Header.Get("Rblx-Challenge-Type") == "twostepverification" {
				rblxChallengeID := resp.Header.Get("Rblx-Challenge-Id")
				rblxChallengeMetadata := resp.Header.Get("Rblx-Challenge-Metadata")
				solveTFAChallengeRes, err = r.SolveTFAChallenge(rblxChallengeID, rblxChallengeMetadata)
				if err != nil {
					return nil, fmt.Errorf("error solving tfa challenge: %w", err)
				}
				continue
			}
		}
		return resp, nil
	}
}
