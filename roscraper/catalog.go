package roscraper

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/kodishim/goblox/robloxapi"
)

func (r *Roscraper) FetchCatalogItemsData(itemIDs ...int) ([]*robloxapi.CatalogItemData, error) {
	if len(itemIDs) > 100 {
		return nil, errors.New("error can't fetch more than 100 item's catalog data at once")
	}
	url := robloxapi.CatalogAPI + "/catalog/items/details"
	type item struct {
		ItemType int `json:"itemType"`
		ID       int `json:"id"`
	}
	var body struct {
		Items []item `json:"items"`
	}
	body.Items = []item{}
	for _, itemID := range itemIDs {
		body.Items = append(body.Items, item{1, itemID})
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error marshaling response body: %w", err)
	}
	resp, err := r.Request(http.MethodPost, url, jsonBody)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	if resp.APIError != nil {
		return nil, resp.APIError
	}
	var respBody robloxapi.FetchCatalogItemsDataResponse
	err = json.Unmarshal(resp.Body, &respBody)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}
	return respBody.ItemsData, nil

}
