package robloxapi

import "fmt"

type APIError struct {
	Code              int               `json:"code"`
	Message           string            `json:"message"`
	UserFacingMessage string            `json:"userFacingMessage"`
	Field             string            `json:"field"`
	FieldData         map[string]string `json:"fieldData"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Message)
}
