package bnm

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// BNM API client
type API struct {
	client *http.Client
}

// Create new API client
func NewAPI(client *http.Client) *API {
	return &API{
		client: client,
	}
}

// Generate URL for BNM API endpoints. `endpoint` can be either a relative or absolute URL.
// If it's a relative URL, it will be appended to "https://api.bnm.gov.my/public/".
// If it's an absolute URL, it will be used as is. Absolute means that it has a non-empty scheme.
func (a *API) formatUrl(endpoint string) (*url.URL, error) {
	parsedUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}

	if parsedUrl.IsAbs() {
		return parsedUrl, nil
	}

	return url.Parse("https://api.bnm.gov.my/public/" + strings.TrimLeft(endpoint, "/"))
}

// Implement `bmn.Requester` interface
func (a *API) Request(endpoint string, queryParams map[string]string, result interface{}) error {
	url, err := a.formatUrl(endpoint)
	if err != nil {
		return fmt.Errorf("error formatting URL: %w", err)
	}

	q := url.Query()
	for key, value := range queryParams {
		q.Add(key, value)
	}
	url.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	// Required by BNM API
	req.Header.Add("Accept", "application/vnd.BNM.API.v1+json")

	response, err := a.client.Do(req)
	if err != nil {
		return fmt.Errorf("error making HTTP request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return nil
}
