package robloxapi

import "time"

type ResaleItemData struct {
	AssetStock         any `json:"assetStock"`
	Sales              int `json:"sales"`
	NumberRemaining    any `json:"numberRemaining"`
	RecentAveragePrice int `json:"recentAveragePrice"`
	OriginalPrice      any `json:"originalPrice"`
	PriceDataPoints    []struct {
		Value int       `json:"value"`
		Date  time.Time `json:"date"`
	} `json:"priceDataPoints"`
	VolumeDataPoints []struct {
		Value int       `json:"value"`
		Date  time.Time `json:"date"`
	} `json:"volumeDataPoints"`
}
