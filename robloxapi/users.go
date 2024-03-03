package robloxapi

import "time"

type AuthenticatedUser struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type User struct {
	ID                     int       `json:"id"`
	Name                   string    `json:"name"`
	DisplayName            string    `json:"displayName"`
	Description            string    `json:"description"`
	Created                time.Time `json:"created"`
	IsBanned               bool      `json:"isBanned"`
	ExternalAppDisplayName any       `json:"externalAppDisplayName"`
	HasVerifiedBadge       bool      `json:"hasVerifiedBadge"`
}

type FetchUsersByUsernameResponse struct {
	Users []*struct {
		RequestedUsername string `json:"requestedUsername"`
		HasVerifiedBadge  bool   `json:"hasVerifiedBadge"`
		ID                int    `json:"id"`
		Name              string `json:"name"`
		DisplayName       string `json:"displayName"`
	} `json:"data"`
}
