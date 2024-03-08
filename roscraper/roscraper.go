package roscraper

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/kodishim/goblox/robloxapi"
)

type Roscraper struct {
	Client       *http.Client
	Proxies      []*url.URL
	CurrentProxy int
	CSRFToken    string
}

func New(timeout int, proxies ...string) (*Roscraper, error) {
	var client http.Client
	var proxyURLs []*url.URL
	client.Timeout = time.Duration(timeout) * time.Second
	if len(proxies) > 0 {
		for _, proxy := range proxies {
			proxyURL, err := url.Parse("http://" + proxy)
			if err != nil {
				return nil, fmt.Errorf("error parsing proxy %s: %w", proxy, err)
			}
			proxyURLs = append(proxyURLs, proxyURL)
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURLs[0])}
	}
	return &Roscraper{
		Client:       &client,
		Proxies:      proxyURLs,
		CurrentProxy: 0,
		CSRFToken:    "",
	}, nil
}

func (r *Roscraper) SwapProxy() bool {
	if len(r.Proxies) <= 1 {
		return false
	}
	if r.CurrentProxy < len(r.Proxies)-1 {
		r.CurrentProxy++
	} else {
		r.CurrentProxy = 0
	}
	r.Client.Transport = &http.Transport{Proxy: http.ProxyURL(r.Proxies[r.CurrentProxy])}
	return true
}

func (r *Roscraper) Request(method string, url string, body []byte) (*robloxapi.Response, error) {
	defer r.SwapProxy()
	for {
		req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
		if err != nil {
			return nil, fmt.Errorf("error forming request: %w", err)
		}
		req.Header = http.Header{
			"Content-Type": {"application/json"},
			"X-Csrf-Token": {r.CSRFToken},
		}
		resp, err := robloxapi.Request(req, r.Client)
		if err != nil {
			return nil, err
		}
		if resp.APIError != nil {
			if resp.APIError.Message == "Token Validation Failed" {
				r.CSRFToken = resp.Header.Get("X-Csrf-Token")
				continue
			}
			if resp.APIError.Message == "Too many requests" || resp.APIError.Message == "InternalServerError" {
				swapped := r.SwapProxy()
				if swapped {
					continue
				} else {
					time.Sleep(time.Second * 10)
					continue
				}
			}
		}
		return resp, nil
	}
}
