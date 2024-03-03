package robloxapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Status     string
	StatusCode int
	Header     http.Header
	Body       []byte
	Request    *http.Request
	APIError   *APIError
}

func (r *Response) populateAPIErr() {
	var apiErr *APIError
	var mapBody map[string]interface{}
	json.Unmarshal(r.Body, &mapBody)
	_, ok := mapBody["message"]
	if ok {
		json.Unmarshal(r.Body, &apiErr)
		r.APIError = apiErr
		return
	}
	_, ok = mapBody["errors"]
	if ok {
		var apiErrs struct {
			Errors []*APIError `json:"errors"`
		}
		json.Unmarshal(r.Body, &apiErrs)
		r.APIError = apiErrs.Errors[0]
	}
}

func NewResponse(httpResponse *http.Response) (*Response, error) {
	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	response := &Response{
		Status:     httpResponse.Status,
		StatusCode: httpResponse.StatusCode,
		Header:     httpResponse.Header,
		Body:       body,
		Request:    httpResponse.Request,
	}
	response.populateAPIErr()
	return response, nil
}

func Request(request *http.Request, client ...*http.Client) (*Response, error) {
	c := http.DefaultClient
	if len(client) > 0 {
		c = client[0]
	}
	resp, err := c.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	response, err := NewResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	return response, nil
}
