package robloxapi

import (
	"time"
)

const (
	TradeStatusInbound   = "Inbound"
	TradeStatusOutbound  = "Outbound"
	TradeStatusCompleted = "Completed"
	TradeStatusInactive  = "Inactive"
)

type TradeSystemMetadata struct {
	MaxItemsPerSide            int     `json:"maxItemsPerSide"`
	MinValueRatio              float64 `json:"minValueRatio"`
	TradeSystemMaxRobuxPercent float64 `json:"tradeSystemMaxRobuxPercent"`
	TradeSystemRobuxFee        float64 `json:"tradeSystemRobuxFee"`
}

type Offer struct {
	UserID       int   `json:"userId"`
	UserAssetIDs []int `json:"userAssetIds"`
	Robux        int   `json:"robux"`
}

func NewOffer(userID int, robux int, userAssetIDs ...int) *Offer {
	return &Offer{UserID: userID, Robux: robux, UserAssetIDs: userAssetIDs}
}

type Trade struct {
	Offers []struct {
		User struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			DisplayName string `json:"displayName"`
		} `json:"user"`
		UserAssets []struct {
			ID                 int    `json:"id"`
			SerialNumber       int    `json:"serialNumber"`
			AssetID            int    `json:"assetId"`
			Name               string `json:"name"`
			RecentAveragePrice int    `json:"recentAveragePrice"`
			OriginalPrice      int    `json:"originalPrice"`
			AssetStock         int    `json:"assetStock"`
			MembershipType     string `json:"membershipType"`
		} `json:"userAssets"`
		Robux int `json:"robux"`
	} `json:"offers"`
	ID   int `json:"id"`
	User struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		DisplayName string `json:"displayName"`
	} `json:"user"`
	Created  time.Time `json:"created"`
	IsActive bool      `json:"isActive"`
	Status   string    `json:"status"`
}

type TradeLog struct {
	ID   int `json:"id"`
	User struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		DisplayName string `json:"displayName"`
	} `json:"user"`
	Created    time.Time `json:"created"`
	Expiration time.Time `json:"expiration"`
	IsActive   bool      `json:"isActive"`
	Status     string    `json:"status"`
}

type FetchTradesResponse struct {
	TradeLogs          []*TradeLog `json:"data"`
	PreviousPageCursor string      `json:"previousPageCursor"`
	NextPageCursor     string      `json:"nextPageCursor"`
}

type SendTradeResponse struct {
	TradeID int `json:"id"`
}

type FetchCanTradeWithResponse struct {
	CanTrade bool   `json:"canTrade"`
	Status   string `json:"status"`
}
