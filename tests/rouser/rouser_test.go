package rouser_test

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kodishim/goblox/rouser"
)

var (
	ruser  *rouser.Rouser
	cookie string
	secret string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	cookie = os.Getenv("COOKIE")
	secret = os.Getenv("SECRET")
	roUser, err := rouser.New(cookie, secret)
	if err != nil {
		log.Fatalf("error creating new rouser: %s", err)
	}
	ruser = roUser
}
