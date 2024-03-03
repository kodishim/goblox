package roscraper_test

import "testing"

func TestFetchUser(t *testing.T) {
	user, err := rscraper.FetchUser(1)
	if err != nil {
		t.Fatalf("error fetching user: %s", err)
	}
	if user.Name != "Roblox" {
		t.Fatalf("expected user.Name to be Roblox, got %s", user.Name)
	}
}

func TestFetchUsersByUsername(t *testing.T) {
	res, err := rscraper.FetchUsersByUsername(false, "Roblox", "builderman")
	if err != nil {
		t.Fatalf("error fetching user by username: %s", err)
	}
	if len(res.Users) != 2 {
		t.Fatalf("expected 2 users to be returned, got %d", len(res.Users))
	}
}
