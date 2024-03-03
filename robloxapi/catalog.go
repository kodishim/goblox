package robloxapi

import "time"

type CatalogItemData struct {
	ID           int    `json:"id"`
	ItemType     string `json:"itemType"`
	AssetType    int    `json:"assetType"`
	BundleType   int    `json:"bundleType"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	ProductID    int    `json:"productId"`
	Genres       []int  `json:"genres"`
	BundledItems []struct {
		Owned bool   `json:"owned"`
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Type  string `json:"type"`
	} `json:"bundledItems"`
	ItemStatus              []int    `json:"itemStatus"`
	ItemRestrictions        []string `json:"itemRestrictions"`
	CreatorHasVerifiedBadge bool     `json:"creatorHasVerifiedBadge"`
	CreatorType             string   `json:"creatorType"`
	CreatorTargetID         int      `json:"creatorTargetId"`
	CreatorName             string   `json:"creatorName"`
	Price                   int      `json:"price"`
	PremiumPricing          struct {
		PremiumDiscountPercentage int `json:"premiumDiscountPercentage"`
		PremiumPriceInRobux       int `json:"premiumPriceInRobux"`
	} `json:"premiumPricing"`
	LowestPrice                  int       `json:"lowestPrice"`
	LowestResalePrice            int       `json:"lowestResalePrice"`
	PriceStatus                  string    `json:"priceStatus"`
	UnitsAvailableForConsumption int       `json:"unitsAvailableForConsumption"`
	PurchaseCount                int       `json:"purchaseCount"`
	FavoriteCount                int       `json:"favoriteCount"`
	OffSaleDeadline              time.Time `json:"offSaleDeadline"`
	CollectibleItemID            string    `json:"collectibleItemId"`
	TotalQuantity                int       `json:"totalQuantity"`
	SaleLocationType             string    `json:"saleLocationType"`
	HasResellers                 bool      `json:"hasResellers"`
	IsOffSale                    bool      `json:"isOffSale"`
	QuantityLimitPerUser         int       `json:"quantityLimitPerUser"`
}

type FetchCatalogItemsDataResponse struct {
	ItemsData []*CatalogItemData `json:"data"`
}
