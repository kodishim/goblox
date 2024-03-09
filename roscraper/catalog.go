package roscraper

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kodishim/goblox/robloxapi"
)

func (r *Roscraper) FetchCatalogItemData(itemID int) (*robloxapi.CatalogItemData, error) {
	itemsData, err := r.FetchCatalogItemsData(itemID)
	if err != nil {
		return nil, err
	}
	if len(itemsData) != 1 {
		return nil, fmt.Errorf("expected 1 items data but got %d", len(itemsData))
	}
	return itemsData[0], nil
}

func (r *Roscraper) FetchCatalogItemsData(itemIDs ...int) ([]*robloxapi.CatalogItemData, error) {
	catalogItemsData := []*robloxapi.CatalogItemData{}
	batches := [][]int{}
	for i := 0; i < len(itemIDs); i += 100 {
		end := i + 100
		if end > len(itemIDs) {
			end = len(itemIDs)
		}
		batches = append(batches, itemIDs[i:end])
	}
	for _, batch := range batches {
		url := robloxapi.CatalogAPI + "/catalog/items/details"
		type item struct {
			ItemType int `json:"itemType"`
			ID       int `json:"id"`
		}
		var body struct {
			Items []item `json:"items"`
		}
		for _, itemID := range batch {
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
		catalogItemsData = append(catalogItemsData, respBody.ItemsData...)
	}
	return catalogItemsData, nil
}
