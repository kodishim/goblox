package roscraper_test

import "testing"

func TestFetchResaleItemData(t *testing.T) {
	data, err := rscraper.FetchResaleItemData(9254254)
	if err != nil {
		t.Fatalf("error fetching resale item data: %s", err)
	}
	if data.Sales <= 0 {
		t.Fatal("got 0 sales, expected more than 0 sales")
	}
}
