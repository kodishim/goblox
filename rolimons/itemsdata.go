package rolimons

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type ItemData struct {
	ID        int
	Name      string
	Acronym   string
	RAP       int
	Value     int
	Projected bool
	Demand    int
	Trend     int
	Rare      bool
}

var ItemsDataCache struct {
	ItemsData   map[int]*ItemData
	LastUpdated time.Time
}

var (
	DemandUnknown  = -1
	DemandTerrible = 0
	DemandLow      = 1
	DemandNormal   = 2
	DemandHigh     = 3
	DemandAmazing  = 4
)

var updateCacheEvery = time.Minute * 10

func FetchItemsData(forceLive ...bool) (map[int]*ItemData, error) {
	useCache := true
	if len(forceLive) > 0 {
		useCache = !forceLive[0]
	}
	if time.Since(ItemsDataCache.LastUpdated) < (updateCacheEvery) && useCache && ItemsDataCache.ItemsData != nil {
		return ItemsDataCache.ItemsData, nil
	}
	itemsData := make(map[int]*ItemData)
	resp, err := http.Get("https://www.rolimons.com/itemapi/itemdetails")
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()
	jsonBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error unexpected response: %d %s", resp.StatusCode, string(jsonBody))
	}
	var body struct {
		Success bool                  `json:"success"`
		Items   map[int][]interface{} `json:"items"`
	}
	err = json.Unmarshal(jsonBody, &body)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}
	for id, details := range body.Items {
		projected := false
		rare := false
		if int(details[7].(float64)) == 1 {
			projected = true
		}
		if int(details[9].(float64)) == 1 {
			rare = true
		}
		rap := int(details[2].(float64))
		value := int(details[3].(float64))
		if value == -1 {
			value = rap
		}
		data := ItemData{
			ID:        id,
			Name:      strings.TrimPrefix(details[0].(string), " "),
			Acronym:   strings.TrimPrefix(details[1].(string), " "),
			RAP:       rap,
			Value:     value,
			Demand:    int(details[5].(float64)),
			Trend:     int(details[6].(float64)),
			Projected: projected,
			Rare:      rare,
		}
		itemsData[id] = &data
	}
	ItemsDataCache.ItemsData = itemsData
	ItemsDataCache.LastUpdated = time.Now()
	return itemsData, nil
}
