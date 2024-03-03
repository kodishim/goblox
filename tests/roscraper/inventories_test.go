package roscraper_test

import (
	"testing"
)

func TestFetchUserItems(t *testing.T) {
	inventories, err := rscraper.FetchUserItems(1, 100, "", "")
	if err != nil {
		t.Fatalf("error fetching inventory: %s", err)
	}
	if len(inventories.Items) != 100 {
		t.Fatalf("expected 100 items to be returned, got %d", len(inventories.Items))
	}
}

func TestFetchAllUserItems(t *testing.T) {
	_, err := rscraper.FetchAllUserItems(3308733)
	if err != nil {
		t.Fatalf("error fetching inventory: %s", err)
	}
}
