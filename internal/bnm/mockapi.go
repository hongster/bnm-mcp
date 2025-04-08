package bnm

import "encoding/json"

type MockAPI struct {
	Endpoint       string
	QueryParams    map[string]string
	ResultJSONText string
}

func (m *MockAPI) Request(endpoint string, queryParams map[string]string, result interface{}) error {
	m.Endpoint = endpoint
	m.QueryParams = queryParams

	err := json.Unmarshal([]byte(m.ResultJSONText), result)
	if err != nil {
		return err
	}

	return nil
}
