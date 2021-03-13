package gqltest

type Variables = map[string]interface{}

type RequestBody struct {
	Query     string    `json:"query"`
	Variables Variables `json:"variables,omitempty"`
}
