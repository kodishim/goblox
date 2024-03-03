package robloxapi

type UserItem struct {
	UserAssetID                int    `json:"userAssetId"`
	SerialNumber               int    `json:"serialNumber"`
	AssetID                    int    `json:"assetId"`
	Name                       string `json:"name"`
	RecentAveragePrice         int    `json:"recentAveragePrice"`
	OriginalPrice              int    `json:"originalPrice"`
	AssetStock                 int    `json:"assetStock"`
	BuildersClubMembershipType int    `json:"buildersClubMembershipType"`
	IsOnHold                   bool   `json:"isOnHold"`
}

type FetchUserItemsResponse struct {
	PreviousPageCursor string      `json:"previousPageCursor"`
	NextPageCursor     string      `json:"nextPageCursor"`
	Items              []*UserItem `json:"data"`
}
