package rouser_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"
	"time"

	"github.com/kodishim/goblox/robloxapi"
)

var tradeData struct {
	TestTrade struct {
		Offers []robloxapi.Offer `json:"offers"`
	} `json:"testTrade"`
}

func init() {
	tradeFile, err := os.ReadFile("data/trades.json")
	if err != nil {
		log.Fatalf("error reading data/trades.json: %s", err)
	}
	err = json.Unmarshal(tradeFile, &tradeData)
	if err != nil {
		log.Fatalf("error unmarshaling data/trades.json: %s", err)
	}
}

func TestSendFetchDeclineTrade(t *testing.T) {
	tradeID, err := ruser.SendTrade([2]robloxapi.Offer(tradeData.TestTrade.Offers))
	if err != nil {
		t.Fatalf("error sending trade: %s", err)
	}
	if tradeID == 0 {
		t.Fatal("trade id is 0")
	}
	trade, err := ruser.FetchTrade(tradeID)
	if err != nil {
		t.Fatalf("error fetching trade: %s", err)
	}
	if trade.ID != tradeID {
		t.Fatalf("sent trade id is %d and fetched trade is %d", tradeID, trade.ID)
	}
	err = ruser.DeclineTrade(tradeID)
	if err != nil {
		t.Fatalf("error declining trade: %s", err)
	}
	time.Sleep(time.Second)
	trade, err = ruser.FetchTrade(tradeID)
	if err != nil {
		t.Fatalf("error fetching trade after declining: %s", err)
	}
	if trade.Status != "Declined" {
		t.Fatalf("expected Declined trade status, got %s", trade.Status)
	}
}

func TestFetchTrades(t *testing.T) {
	resp, err := ruser.FetchTrades(robloxapi.TradeStatusInactive, robloxapi.Limit100, "", "")
	if err != nil {
		t.Fatalf("error fetching trades %s", err)
	}
	if len(resp.TradeLogs) <= 0 {
		t.Fatal("0 trades returned")
	}
}

func TestFetchAllTrades(t *testing.T) {
	_, err := ruser.FetchAllTrades(robloxapi.TradeStatusOutbound)
	if err != nil {
		t.Fatalf("error fetching trades %s", err)
	}
}

func TestFetchTradeSystemMetata(t *testing.T) {
	tradeSystemMetadata, err := ruser.FetchTradeSystemMetadata()
	if err != nil {
		t.Fatalf("error fetching trade system metadata %s", err)
	}
	if tradeSystemMetadata.MaxItemsPerSide != 4 {
		t.Fatalf("MaxItemsPerSide is %d, expected 4", tradeSystemMetadata.MaxItemsPerSide)
	}
}

func TestExpireOutdatedInboundTrades(t *testing.T) {
	err := ruser.ExpireOutdatedInboundTrades()
	if err != nil {
		t.Fatalf("error expiring outdated inbound trades: %s", err)
	}
}

func TestFetchCanTradeWith(t *testing.T) {
	canTrade, err := ruser.FetchCanTradeWith(560349464)
	if err != nil {
		t.Fatalf("error fecthing can trade with: %s", err)
	}
	if !canTrade {
		t.Fatalf("canTrade = false for user %d expected true", 560349464)
	}
	canTrade, err = ruser.FetchCanTradeWith(1)
	if err != nil {
		t.Fatalf("error fecthing can trade with: %s", err)
	}
	if canTrade {
		t.Fatalf("canTrade = true for user 1 expected true")
	}
}
