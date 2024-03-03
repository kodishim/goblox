package roscraper

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/robloxapi"
)

func (r *Roscraper) FetchResaleItemData(itemID int) (*robloxapi.ResaleItemData, error) {
	url := robloxapi.EconomyAPI + fmt.Sprintf("/assets/%d/resale-data", itemID)
	resp, err := r.Request(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	if resp.APIError != nil {
		return nil, resp.APIError
	}
	var itemData *robloxapi.ResaleItemData
	err = json.Unmarshal(resp.Body, &itemData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}
	return itemData, nil
}
