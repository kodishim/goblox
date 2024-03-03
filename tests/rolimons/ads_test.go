package rolimons_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"example.com/rolimons"
	"github.com/joho/godotenv"
)

var ruser *rolimons.RolimonsUser
var cookie string

var rolimonsData struct {
	UserID      int `json:"userId"`
	TestTradeAd struct {
		Offer       []int    `json:"offer"`
		Request     []int    `json:"request"`
		RequestTags []string `json:"requestTags"`
	} `json:"testTradeAd"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	cookie = os.Getenv("COOKIE")
	ruser = rolimons.New(cookie)
	rolimonsFile, err := os.ReadFile("data/rolimons.json")
	if err != nil {
		log.Fatalf("error reading data/rolimons.json: %s", err)
	}
	err = json.Unmarshal(rolimonsFile, &rolimonsData)
	if err != nil {
		log.Fatalf("error unmarshaling data/rolimons.json: %s", err)
	}
}

func TestCreateAD(t *testing.T) {
	err := ruser.CreateAD(rolimonsData.UserID, rolimonsData.TestTradeAd.Offer, rolimonsData.TestTradeAd.Request, rolimonsData.TestTradeAd.RequestTags)
	if err != nil && err != rolimons.ErrCooldown {
		t.Fatalf("error creating trade ad: %s", err)
	}
}
