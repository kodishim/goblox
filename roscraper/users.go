package roscraper

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/robloxapi"
)

func (r *Roscraper) FetchUser(userID int) (*robloxapi.User, error) {
	url := robloxapi.UsersAPI + fmt.Sprintf("/users/%d", userID)
	resp, err := r.Request(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	if resp.APIError != nil {
		return nil, resp.APIError
	}
	var user robloxapi.User
	err = json.Unmarshal(resp.Body, &user)
	if err != nil {
		return nil, fmt.Errorf("error fetching user: %w", err)
	}
	return &user, nil
}

func (r *Roscraper) FetchUsersByUsername(excludeBannedUsers bool, usernames ...string) (*robloxapi.FetchUsersByUsernameResponse, error) {
	url := robloxapi.UsersAPI + "/usernames/users"
	var body struct {
		Usernames          []string `json:"usernames"`
		ExcludeBannedUsers bool     `json:"excludeBannedUsers"`
	}
	body.Usernames = usernames
	body.ExcludeBannedUsers = excludeBannedUsers
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}
	resp, err := r.Request(http.MethodPost, url, jsonBody)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	if resp.APIError != nil {
		return nil, resp.APIError
	}
	var respBody robloxapi.FetchUsersByUsernameResponse
	err = json.Unmarshal(resp.Body, &respBody)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}
	return &respBody, nil
}
