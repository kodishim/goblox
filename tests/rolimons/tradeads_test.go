package rolimons

import (
	"testing"

	"github.com/kodishim/goblox/rolimons"
)

func TestFetchRecentTradeADUsers(t *testing.T) {
	ids, err := rolimons.FetchRecentTradeADUsers()
	if err != nil {
		t.Fatalf("error fetching items data (1st attempt): %s", err)
	}
	if len(ids) < 10 {
		t.Fatalf("only got %d user IDs", len(ids))
	}
}
