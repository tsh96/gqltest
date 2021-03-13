package gqltest

import (
	"encoding/json"
)

type ResponseBody struct {
	Errors []interface{}   `json:"errors,omitempty"`
	Data   json.RawMessage `json:"data,omitempty"`
}
