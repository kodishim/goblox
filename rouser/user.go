package rouser

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kodishim/goblox/robloxapi"
)

func (r *Rouser) FetchAuthenticatedUser() (*robloxapi.AuthenticatedUser, error) {
	url := robloxapi.UsersAPI + "/users/authenticated"
	resp, err := r.Request(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	if resp.APIError != nil {
		return nil, resp.APIError
	}
	var authenticatedUser robloxapi.AuthenticatedUser
	err = json.Unmarshal(resp.Body, &authenticatedUser)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}
	return &authenticatedUser, nil
}
