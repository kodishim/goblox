package rouser

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/kodishim/goblox/robloxapi"
	"github.com/xlzd/gotp"
)

type SolveTFAChallengeRes struct {
	RblxChallengeId       string
	RblxChallengeMetadata string
	RblxChallengeType     string
}

func (r *Rouser) SolveTFAChallenge(rblxChallengeID string, rblxChallengeMetadata string) (*SolveTFAChallengeRes, error) {
	rblxChallengeMetadataBytes, err := base64.StdEncoding.DecodeString(rblxChallengeMetadata)
	if err != nil {
		return nil, fmt.Errorf("error decoding rblx challenge metadata %s: %w", rblxChallengeMetadata, err)
	}
	var rChallengeMetadata robloxapi.RblxChallengeMetadata
	err = json.Unmarshal(rblxChallengeMetadataBytes, &rChallengeMetadata)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling rblx challenge metadata: %w", err)
	}
	verificationToken, err := r.VerifyTFAChallenge(&rChallengeMetadata)
	if err != nil {
		return nil, fmt.Errorf("error verifying TFA Challenge: %w", err)
	}
	solvedChallengeMetadata := robloxapi.NewSolvedChallengeMetadata(verificationToken, rChallengeMetadata.ShouldShowRememberDeviceCheckbox, rChallengeMetadata.ChallengeID, rChallengeMetadata.ActionType)
	err = r.ContinueTFAChallenge(rblxChallengeID, solvedChallengeMetadata)
	if err != nil {
		return nil, fmt.Errorf("error continuing TFA Challenge: %w", err)
	}
	jsonSolvedChallengeMetadata, err := json.Marshal(solvedChallengeMetadata)
	if err != nil {
		return nil, fmt.Errorf("error marshaling solved challenge metadata: %w", err)
	}
	base64SolvedChallengeMetadata := base64.RawStdEncoding.EncodeToString(jsonSolvedChallengeMetadata)
	return &SolveTFAChallengeRes{rblxChallengeID, base64SolvedChallengeMetadata, "twostepverification"}, nil
}

func (r *Rouser) VerifyTFAChallenge(rblxChallengeMetadata *robloxapi.RblxChallengeMetadata) (string, error) {

	url := robloxapi.TwoStepVerificationAPI + fmt.Sprintf("/users/%d/challenges/authenticator/verify", r.User.ID)
	var body struct {
		ChallengeID string `json:"challengeId"`
		ActionType  string `json:"actionType"`
		Code        string `json:"code"`
	}
	body.ChallengeID = rblxChallengeMetadata.ChallengeID
	body.ActionType = rblxChallengeMetadata.ActionType
	for {
		body.Code = gotp.NewDefaultTOTP(r.TFASecret).Now()
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return "", fmt.Errorf("error marshaling request body: %w", err)
		}
		resp, err := r.Request(http.MethodPost, url, jsonBody)
		if err != nil {
			return "", fmt.Errorf("error making request: %w", err)
		}
		if resp.APIError != nil {
			if resp.APIError.Message == "Authenticator code already used." {
				time.Sleep(time.Second * 30)
				continue
			}
			return "", resp.APIError
		}
		var successBody struct {
			VerificationToken string `json:"verificationToken"`
		}
		err = json.Unmarshal(resp.Body, &successBody)
		if err != nil {
			return "", fmt.Errorf("error unmarshaling response body: %w", err)
		}
		return successBody.VerificationToken, nil
	}
}

func (r *Rouser) ContinueTFAChallenge(rblxChallengeID string, solvedChallengeMetadata *robloxapi.SolvedChallengeMetadata) error {
	url := "https://apis.roblox.com/challenge/v1/continue"
	solvedChallengeMetadataJson, err := json.Marshal(solvedChallengeMetadata)
	if err != nil {
		return fmt.Errorf("error marshaling solved challenge metadata: %w", err)
	}
	var body struct {
		ChallengeID       string `json:"challengeId"`
		ChallengeType     string `json:"challengeType"`
		ChallengeMetadata string `json:"challengeMetadata"`
	}
	body.ChallengeID = rblxChallengeID
	body.ChallengeType = "twostepverification"
	body.ChallengeMetadata = string(solvedChallengeMetadataJson)
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error marshaling json body: %w", err)
	}
	resp, err := r.Request(http.MethodPost, url, jsonBody)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	if resp.APIError != nil {
		return resp.APIError
	}
	return nil
}
