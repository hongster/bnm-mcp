package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFormUrl(t *testing.T) {
	tests := []struct {
		name     string
		endpoint string
		want     string
	}{
		{
			name:     "basic endpoint",
			endpoint: "test",
			want:     "https://api.bnm.gov.my/public/test",
		},
		{
			name:     "endpoint with leading slash",
			endpoint: "/test",
			want:     "https://api.bnm.gov.my/public/test",
		},
		{
			name:     "endpoint with multiple leading slashes",
			endpoint: "///test",
			want:     "https://api.bnm.gov.my/public/test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formUrl(tt.endpoint); got != tt.want {
				t.Errorf("formUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

// FIXME
func TestRequest(t *testing.T) {
	type testResponse struct {
		Message string `json:"message"`
	}

	tests := []struct {
		name           string
		endpoint       string
		queryParams    map[string]string
		serverResponse string
		statusCode     int
		wantErr        bool
		wantResponse   *testResponse
	}{
		{
			name:           "successful request",
			queryParams:    map[string]string{"key": "value"},
			serverResponse: `{"message": "success"}`,
			statusCode:     http.StatusOK,
			wantErr:        false,
			wantResponse:   &testResponse{Message: "success"},
		},
		{
			name:           "server error",
			queryParams:    nil,
			serverResponse: `{"message": "error"}`,
			statusCode:     http.StatusInternalServerError,
			wantErr:        true,
			wantResponse:   &testResponse{},
		},
		{
			name:           "invalid JSON response",
			queryParams:    nil,
			serverResponse: `invalid json`,
			statusCode:     http.StatusOK,
			wantErr:        true,
			wantResponse:   &testResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Check Accept header
				if r.Header.Get("Accept") != "application/vnd.BNM.API.v1+json" {
					t.Error("Missing or incorrect Accept header")
				}

				// Check query parameters
				if tt.queryParams != nil {
					for key, want := range tt.queryParams {
						if got := r.URL.Query().Get(key); got != want {
							t.Errorf("Query param %s = %v, want %v", key, got, want)
						}
					}
				}

				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.serverResponse))
			}))
			defer server.Close()

			// Create a test response struct
			result := &testResponse{}

			// Call the function
			err := Request(server.URL, tt.queryParams, result)

			// Check error
			if (err != nil) != tt.wantErr {
				t.Errorf("Request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// If we don't expect an error, check the response
			if !tt.wantErr && result.Message != tt.wantResponse.Message {
				t.Errorf("Request() got = %v, want %v", result, tt.wantResponse)
			}
		})
	}
}
