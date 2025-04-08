package bnm

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFormatUrl(t *testing.T) {
	api := NewAPI(http.DefaultClient)

	tests := []struct {
		name     string
		endpoint string
		want     string
		wantErr  bool
	}{
		{
			name:     "relative URL",
			endpoint: "exchange-rate",
			want:     "https://api.bnm.gov.my/public/exchange-rate",
			wantErr:  false,
		},
		{
			name:     "relative URL with leading slash",
			endpoint: "/exchange-rate",
			want:     "https://api.bnm.gov.my/public/exchange-rate",
			wantErr:  false,
		},
		{
			name:     "absolute URL",
			endpoint: "https://example.com/api",
			want:     "https://example.com/api",
			wantErr:  false,
		},
		{
			name:     "invalid URL",
			endpoint: "://invalid",
			want:     "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := api.formatUrl(tt.endpoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("formatUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got.String() != tt.want {
				t.Errorf("formatUrl() = %v, want %v", got.String(), tt.want)
			}
		})
	}
}

func TestRequest(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check headers
		if r.Header.Get("Accept") != "application/vnd.BNM.API.v1+json" {
			t.Error("Expected Accept header not found")
		}

		// Test different endpoints and responses
		switch r.URL.Path {
		case "/test":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"data": "test"}`))
		case "/error":
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	// Create API client with test server
	api := NewAPI(&http.Client{})

	tests := []struct {
		name        string
		endpoint    string
		queryParams map[string]string
		wantErr     bool
	}{
		{
			name:        "successful request",
			endpoint:    server.URL + "/test",
			queryParams: map[string]string{"param": "value"},
			wantErr:     false,
		},
		{
			name:        "bad request",
			endpoint:    server.URL + "/error",
			queryParams: nil,
			wantErr:     true,
		},
		{
			name:        "not found",
			endpoint:    server.URL + "/notfound",
			queryParams: nil,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result map[string]interface{}
			err := api.Request(tt.endpoint, tt.queryParams, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("Request() error = %v, endpoint %v", err, tt.endpoint)
			}
		})
	}
}
