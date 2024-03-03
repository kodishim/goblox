package rolimons_test

import (
	"testing"
	"time"

	"example.com/rolimons"
)

func init() {}

func TestFetchItemsData(t *testing.T) {
	start := time.Now()
	_, err := rolimons.FetchItemsData()
	if err != nil {
		t.Fatalf("error fetching items data (1st attempt): %s", err)
	}
	firstTime := time.Since(start)
	start = time.Now()
	items, err := rolimons.FetchItemsData()
	if err != nil {
		t.Fatalf("error fetching items data (2st attempt): %s", err)
	}
	secondTime := time.Since(start)
	if firstTime < secondTime {
		t.Fatalf("request took longer without cache than with cache, first time: %d ms, second time: %d ms", firstTime.Milliseconds(), secondTime.Milliseconds())
	}
	if items[362081769] == nil {
		t.Fatal("No virtual commando in response")
	}
	if items[362081769].Name != "Virtual Commando" {
		t.Fatalf("Virtual Commando had name %s", items[362081769].Name)
	}
	if items[362081769].Acronym != "VC" {
		t.Fatalf("Virtual Commando had acronym %s", items[362081769].Acronym)
	}
}
