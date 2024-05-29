package rolimons

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type RolimonsUser struct {
	Cookie string
}

func New(cookie string) *RolimonsUser {
	cookie = "_RoliVerification=" + cookie
	return &RolimonsUser{cookie}
}

var (
	TagDemand = "demand"
	TagAny    = "any"
	TagRares  = "rares"
	TagRAP    = "rap"
)

var (
	ErrCooldown              = errors.New("cooldown")
	ErrUnauthorized          = errors.New("unauthorized")
	ErrUserNoLongerOwnsItems = errors.New("user_no_longer_owns_items")
)

func (u *RolimonsUser) CreateAD(userID int, offer []int, request []int, tags []string) error {
	url := "https://api.rolimons.com/tradeads/v1/createad"

	var body struct {
		PlayerID       int      `json:"player_id"`
		OfferItemIDs   []int    `json:"offer_item_ids"`
		RequestItemIDs []int    `json:"request_item_ids"`
		RequestTags    []string `json:"request_tags"`
	}
	body.PlayerID = userID
	body.OfferItemIDs = offer
	body.RequestItemIDs = request
	body.RequestTags = tags
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error marshaling request body: %w", err)
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	req.Header.Set("Cookie", u.Cookie)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}
	if strings.Contains(string(respBody), "Ad creation cooldown has not elapsed") {
		return ErrCooldown
	}
	if strings.Contains(string(respBody), "Verification error") {
		return ErrUnauthorized
	}
	if strings.Contains(string(respBody), "Player does not own all offered items") {
		return ErrUserNoLongerOwnsItems
	}
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return fmt.Errorf("unexpected response: %d %s", resp.StatusCode, string(respBody))
	}
	return nil
}
