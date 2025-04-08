package bnm

// Interface to be implemented by API client
type Requester interface {
	// Request sends a GET request to the given endpoint with the given query parameters
	// and unmarshals the response into the given result interface.
	// `endpoint` - (URL) API endpoint.
	// `params` - Query parameters.
	// `result` - Result of request (decoded JSON).
	Request(endpoint string, queryParams map[string]string, result interface{}) error
}
