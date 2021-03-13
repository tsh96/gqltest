package gqltest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
)

func NewRequest(query string, options ...*RequestOption) (*http.Request, error) {
	body := &RequestBody{
		Query: query,
	}

	method := "POST"
	target := "/"
	for _, option := range options {
		if option.Method != "" {
			method = option.Method
		}
		if option.Target != "" {
			target = option.Target
		}
		if option.Variables != nil {
			body.Variables = option.Variables
		}
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	r := httptest.NewRequest(method, target, bytes.NewReader(bodyBytes))
	r.Header.Add("content-type", "application/json")

	return r, nil
}

func NewRequestFromFile(filename string, options ...*RequestOption) (*http.Request, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return NewRequest(string(fileBytes), options...)
}
