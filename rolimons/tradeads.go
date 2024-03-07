package rolimons

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Define the structure of the response based on the provided example
type TradeAd struct {
	UserID int `json:"user_id"`
}

type TradeAdsResponse struct {
	Success  bool            `json:"success"`
	TradeAds [][]interface{} `json:"trade_ads"`
}

func FetchRecentTradeADUsers() ([]int, error) {
	url := "https://api.rolimons.com/tradeads/v1/getrecentads"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	var successBody TradeAdsResponse
	err = json.Unmarshal(respBody, &successBody)
	if err != nil {
		return nil, fmt.Errorf("error unmmarshaling response body: %w", err)
	}
	if !successBody.Success {
		return nil, fmt.Errorf("error unexpected response: %d %s", resp.StatusCode, string(respBody))
	}
	var userIDs []int
	for _, ad := range successBody.TradeAds {
		if len(ad) > 2 {
			userID, ok := ad[2].(float64)
			if ok {
				userIDs = append(userIDs, int(userID))
			}
		}
	}
	return userIDs, nil
}
