package roscraper_test

import "testing"

func TestFetchCatalogItemsData(t *testing.T) {
	data, err := rscraper.FetchCatalogItemsData(152980783, 152980589)
	if err != nil {
		t.Fatalf("error fetching catalog items data: %s", err)
	}
	if len(data) != 2 {
		t.Fatalf("expected 2 items to be returned, got %d", len(data))
	}
}
