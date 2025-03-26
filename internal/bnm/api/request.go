package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Generate full API URL with the given suffix (`endpoint`).
// For example, given "exchange-rate", returns "https://api.bnm.gov.my/public/exchange-rate".
func formUrl(endpoint string) string {
	return "https://api.bnm.gov.my/public/" + strings.TrimLeft(endpoint, "/")
}

// Perform a HTTP GET request for JSON data.
//
// Arguments:
//   - endpoint: Request URL. Can be a full URL (e.g. "https://api.bnm.gov.my/public/consumer-alert")
//     or a relative URL. If relative URL is given, it will be appended "https://api.bnm.gov.my/public/".
//   - queryParams: A map of key-value pairs for the query string parameters.
//   - result: A pointer to a struct to decode the JSON response into.
//
// Returns:
//   - An error, if any occurred.
func Request(endpoint string, queryParams map[string]string, result interface{}) error {
	var u *url.URL
	var err error
	if endpoint[0:4] == "http" {
		u, err = url.Parse(endpoint)
	} else {
		// If this is a relative URL, append the base URL
		u, err = url.Parse(formUrl(endpoint))
	}
	if err != nil {
		return fmt.Errorf("error parsing URL: %w", err)
	}

	q := u.Query()
	for key, value := range queryParams {
		q.Add(key, value)
	}
	u.RawQuery = q.Encode()

	client := &http.Client{}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	// Required by BNM API
	req.Header.Add("Accept", "application/vnd.BNM.API.v1+json")

	response, err := client.Do(req)
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
