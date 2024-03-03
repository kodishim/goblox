package roscraper

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kodishim/goblox/robloxapi"
)

func (r *Roscraper) FetchUserItems(userID int, limit int, cursor string, sortOrder string) (*robloxapi.FetchUserItemsResponse, error) {
	url := robloxapi.InventoryAPI + fmt.Sprintf("/users/%d/assets/collectibles?limit=%d", userID, limit)
	if cursor != "" {
		url += fmt.Sprintf("&cursor=%s", cursor)
	}
	if sortOrder != "" {
		url += fmt.Sprintf("&sortOrder=%s", sortOrder)
	}
	resp, err := r.Request(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	if resp.APIError != nil {
		return nil, resp.APIError
	}
	var respBody robloxapi.FetchUserItemsResponse
	err = json.Unmarshal(resp.Body, &respBody)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}
	return &respBody, nil
}

func (r *Roscraper) FetchAllUserItems(userID int) ([]*robloxapi.UserItem, error) {
	allUserItems := []*robloxapi.UserItem{}
	cur := ""
	for {
		resp, err := r.FetchUserItems(userID, 100, cur, "")
		if err != nil {
			return nil, err
		}
		cur = resp.NextPageCursor
		if cur == "" {
			break
		}
		allUserItems = append(allUserItems, resp.Items...)
	}
	return allUserItems, nil
}
