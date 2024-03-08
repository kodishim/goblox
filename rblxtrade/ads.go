package rblxtrade

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var (
	ErrCooldown     = errors.New("cooldown")
	ErrUnauthorized = errors.New("unauthorized")
)

type RblxTradeUser struct {
	Cookie        string
	CSRFToken     string
	RblxCSRFToken string
}

func New(cookie string) *RblxTradeUser {
	return &RblxTradeUser{cookie, "", ""}
}

func (u *RblxTradeUser) getCookieHeader() string {
	cookie := "rblx-sess-v1=" + u.Cookie
	if u.CSRFToken != "" {
		cookie += "; rblx-csrf-v4=" + u.RblxCSRFToken
	}
	return cookie
}

var (
	TagAny        = 1
	TagRares      = 2
	TagDemand     = 3
	TagRap        = 4 // unsure
	TagRobux      = 4 // unsure
	TagProjecteds = 5 // unsure
	TagUpgrade    = 8
	TagDowngrade  = 9
	TagValue      = 0 // unsure
	TagAdds       = 0 // unsure
)

func (u *RblxTradeUser) CreateAD(userID int, offer []int, request []int, offerRobux int, requestRobux int, requestTags []int) error {
	for {
		url := "https://rblx.trade/api/v2/trade-ads/create"
		type robux struct {
			Amount int `json:"amount"`
			Side   int `json:"side"`
		}
		type item struct {
			AssetId int `json:"assetId"`
			Side    int `json:"side"`
		}
		type tag struct {
			TagId int `json:"tagId"`
			Side  int `json:"side"`
		}
		var body struct {
			Items          []item  `json:"items"`
			AdTags         []tag   `json:"adTags"`
			Robux          []robux `json:"robux"`
			RequestComment string  `json:"requestComment"`
			OfferComment   string  `json:"offerComment"`
		}
		body.Items = make([]item, 0)
		body.AdTags = make([]tag, 0)
		body.Robux = make([]robux, 0)
		body.RequestComment = ""
		body.OfferComment = ""
		for _, id := range offer {
			body.Items = append(body.Items, item{id, 1})
		}
		for _, id := range request {
			body.Items = append(body.Items, item{id, 2})
		}
		for _, t := range requestTags {
			body.AdTags = append(body.AdTags, tag{t, 2})
		}
		if offerRobux != 0 {
			body.Robux = append(body.Robux, robux{offerRobux, 1})
		}
		if requestRobux != 0 {
			body.Robux = append(body.Robux, robux{requestRobux, 2})
		}
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("error marshaling json body: %w", err)
		}
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
		if err != nil {
			return fmt.Errorf("error forming request: %w", err)
		}
		req.Header.Set("Cookie", u.getCookieHeader())
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Csrf-Token", u.CSRFToken)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("error sending request: %w", err)
		}
		defer resp.Body.Close()
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error reading response body: %w", err)
		}
		if strings.Contains(string(respBody), "LoginRequired") {
			return ErrUnauthorized
		}
		if strings.Contains(string(respBody), "TradeAdCooldown") {
			return ErrCooldown
		}
		if strings.Contains(string(respBody), "CsrfValidationFailed") {
			u.CSRFToken = resp.Header.Get("X-Csrf-Token")
			cookies := strings.Split(resp.Header.Get("Set-Cookie"), ";")
			for _, cookie := range cookies {
				cookie = strings.TrimSpace(cookie)
				if strings.HasPrefix(cookie, "rblx-csrf-v4") {
					parts := strings.SplitN(cookie, "=", 2)
					if len(parts) == 2 {
						u.RblxCSRFToken = parts[1]
					}
					break
				}
			}
			continue
		}
		if resp.StatusCode != 200 && resp.StatusCode != 201 {
			return fmt.Errorf("unexpected response: %d %s", resp.StatusCode, string(respBody))
		}
		return nil
	}
}
