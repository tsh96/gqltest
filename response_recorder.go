package gqltest

import (
	"encoding/json"
	"net/http/httptest"
)

type ResponseRecorder struct {
	*httptest.ResponseRecorder
}

func NewRecorder() *ResponseRecorder {
	return &ResponseRecorder{
		httptest.NewRecorder(),
	}
}

func (r *ResponseRecorder) ResponseBody() (*ResponseBody, error) {
	body := &ResponseBody{}
	err := json.Unmarshal(r.Body.Bytes(), body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
