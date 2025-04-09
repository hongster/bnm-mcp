package exchangerate

import (
	"testing"

	"github.com/hongster/bnm-mcp/internal/bnm"
)

func TestRequest(t *testing.T) {
	api := bnm.MockAPI{
		ResultJSONText: `{
  "data": [
    {
      "currency_code": "INR",
      "unit": 100,
      "rate": {
        "date": "2025-04-09",
        "buying_rate": 5.2145,
        "selling_rate": 5.2226,
        "middle_rate": 5.2186
      }
    },
    {
      "currency_code": "SGD",
      "unit": 1,
      "rate": {
        "date": "2025-04-09",
        "buying_rate": 3.3284,
        "selling_rate": 3.3331,
        "middle_rate": 3.3307
      }
    }
  ],
  "meta": {
    "quote": "rm",
    "session": "0900",
    "last_updated": "2025-04-09 18:21:18",
    "total_result": 2
  }
}`,
	}

	currencies, err := Request(&api)
	t.Run("request", func(t *testing.T) {
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	})

	t.Run("endpoint", func(t *testing.T) {
		if api.Endpoint != "exchange-rate" {
			t.Errorf("unexpected endpoint: %s", api.Endpoint)
		}
	})

	t.Run("currencies", func(t *testing.T) {
		if len(currencies) != 2 {
			t.Errorf("unexpected result length: %d", len(currencies))
		}
	})
}
