package rouser

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kodishim/goblox/robloxapi"
)

// FetchTrade gets detailed information about a trade.
func (r *Rouser) FetchTrade(tradeID int) (*robloxapi.Trade, error) {
	url := robloxapi.TradesAPI + fmt.Sprintf("/trades/%d", tradeID)
	resp, err := r.Request(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	if resp.APIError != nil {
		return nil, resp.APIError
	}
	var trade robloxapi.Trade
	err = json.Unmarshal(resp.Body, &trade)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}
	return &trade, nil
}

// FetchTrades fetches trade logs for all the trades with the specified trade status.
func (r *Rouser) FetchTrades(tradeStatus string, limit int, cursor string, sortOrder string) (*robloxapi.FetchTradesResponse, error) {
	url := robloxapi.TradesAPI + fmt.Sprintf("/trades/%s", tradeStatus)
	url += fmt.Sprintf("?limit=%d", limit)
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
	var respBody robloxapi.FetchTradesResponse
	err = json.Unmarshal(resp.Body, &respBody)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}
	return &respBody, nil
}

// FetchAllTrades fetches all trade logs for all the trades with the specified trade status.
func (r *Rouser) FetchAllTrades(tradeStatus string) ([]*robloxapi.TradeLog, error) {
	cur := ""
	var allTrades []*robloxapi.TradeLog
	for {
		resp, err := r.FetchTrades(tradeStatus, robloxapi.Limit100, cur, "")
		if err != nil {
			return nil, err
		}
		cur = resp.NextPageCursor
		allTrades = append(allTrades, resp.TradeLogs...)
		if cur == "" {
			break
		}
	}
	return allTrades, nil
}

// SendTrade sends a trade.
func (r *Rouser) SendTrade(offers [2]robloxapi.Offer) (int, error) {
	url := robloxapi.TradesAPI + "/trades/send"
	var body struct {
		Offers []robloxapi.Offer `json:"offers"`
	}
	body.Offers = []robloxapi.Offer{offers[0], offers[1]}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return 0, fmt.Errorf("error marshaling request body: %w", err)
	}
	resp, err := r.Request(http.MethodPost, url, jsonBody)
	if err != nil {
		return 0, fmt.Errorf("error sending request: %w", err)
	}
	if resp.APIError != nil {
		return 0, resp.APIError
	}
	var respBody robloxapi.SendTradeResponse
	err = json.Unmarshal(resp.Body, &respBody)
	if err != nil {
		return 0, fmt.Errorf("error unmarshaling response body: %w", err)
	}
	return respBody.TradeID, nil
}

// CounterTrade counters the trade with the given ID.
func (r *Rouser) CounterTrade(tradeID int, offers [2]robloxapi.Offer) (int, error) {
	url := robloxapi.TradesAPI + fmt.Sprintf("/trades/%d/counter", tradeID)
	var body struct {
		Offers []robloxapi.Offer `json:"offers"`
	}
	body.Offers = []robloxapi.Offer{offers[0], offers[1]}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return 0, fmt.Errorf("error marshaling request body: %w", err)
	}
	resp, err := r.Request(http.MethodPost, url, jsonBody)
	if err != nil {
		return 0, fmt.Errorf("error sending request: %w", err)
	}
	if resp.APIError != nil {
		return 0, resp.APIError
	}
	var respBody robloxapi.SendTradeResponse
	err = json.Unmarshal(resp.Body, &respBody)
	if err != nil {
		return 0, fmt.Errorf("error unmarshaling response body: %w", err)
	}
	return respBody.TradeID, nil
}

// AcceptTrade accepts the trade with the given id.
func (r *Rouser) AcceptTrade(tradeID int) error {
	url := robloxapi.TradesAPI + fmt.Sprintf("/trades/%d/accept", tradeID)
	resp, err := r.Request(http.MethodPost, url, nil)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	if resp.APIError != nil {
		return resp.APIError
	}
	return nil
}

// ExpireOutdatedInboundTrades expires all outdated inbound trades.
func (r *Rouser) ExpireOutdatedInboundTrades() error {
	url := robloxapi.TradesAPI + "/trades/expire-outdated"
	resp, err := r.Request(http.MethodPost, url, nil)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	if resp.APIError != nil {
		return resp.APIError
	}
	return nil
}

// FetchCanTradeWith returns true if the user can trade with the specified user.
func (r *Rouser) FetchCanTradeWith(userID int) (bool, error) {
	url := robloxapi.TradesAPI + fmt.Sprintf("/users/%d/can-trade-with", userID)
	resp, err := r.Request(http.MethodGet, url, nil)
	if err != nil {
		return false, fmt.Errorf("error making request: %w", err)
	}
	if resp.APIError != nil {
		return false, resp.APIError
	}
	var respBody robloxapi.FetchCanTradeWithResponse
	err = json.Unmarshal(resp.Body, &respBody)
	if err != nil {
		return false, fmt.Errorf("error unmarshaling response body: %w", err)
	}
	return respBody.CanTrade, nil
}

// DeclineTrade declines the trade with the given ID.
func (r *Rouser) DeclineTrade(tradeID int) error {
	url := robloxapi.TradesAPI + fmt.Sprintf("/trades/%d/decline", tradeID)
	resp, err := r.Request(http.MethodPost, url, nil)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	if resp.APIError != nil {
		return resp.APIError
	}
	return nil
}

// FetchTradeSystemMetadata fetches information about the Roblox trade system.
func (r *Rouser) FetchTradeSystemMetadata() (*robloxapi.TradeSystemMetadata, error) {
	url := robloxapi.TradesAPI + "/trades/metadata"
	resp, err := r.Request(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	if resp.APIError != nil {
		return nil, resp.APIError
	}
	var tradeSystemMetadata robloxapi.TradeSystemMetadata
	err = json.Unmarshal(resp.Body, &tradeSystemMetadata)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}
	return &tradeSystemMetadata, nil
}
