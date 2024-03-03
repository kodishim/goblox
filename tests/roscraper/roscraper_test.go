package roscraper_test

import (
	"log"
	"os"
	"strings"

	"github.com/kodishim/goblox/roscraper"
)

var rscraper *roscraper.Roscraper
var proxies []string

func init() {
	proxiesFile, err := os.ReadFile("data/proxies.txt")
	if err != nil {
		log.Fatalf("error reading data/proxies.txt: %s", err)
	}
	proxies = strings.Split(string(proxiesFile), "\n")
	rscraper, err = roscraper.New(30, proxies...)
	if err != nil {
		log.Fatalf("error creaing roscraper: %s", err)
	}
}
