package consumeralert

import (
	"testing"

	"github.com/hongster/bnm-mcp/internal/bnm"
)

func TestRequest(t *testing.T) {
	api := bnm.MockAPI{
		ResultJSONText: `{
  "data": [
    {
      "name": "1globalcash",
      "regisration_number": "TEST123",
      "added_date": "2012-07-13",
      "websites": []
    },
    {
      "name": "Zenith Gold International Limited (ZGI)",
      "regisration_number": "",
      "added_date": "2016-02-25",
      "websites": [
        "http://www.zenithgolds.com",
        "http://zenithgoldrocks.wordpress.com"
      ]
    }
  ],
  "meta": {
    "last_updated": "2020-07-14 22:32:10",
    "total_result": 2
  }
}`,
	}

	companies, err := Request(&api)
	t.Run("request", func(t *testing.T) {
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	})

	t.Run("endpoint", func(t *testing.T) {
		if api.Endpoint != "consumer-alert" {
			t.Errorf("unexpected endpoint: %s", api.Endpoint)
		}
	})

	t.Run("companies", func(t *testing.T) {
		if len(companies) != 2 {
			t.Errorf("unexpected result length: %d", len(companies))
		}
	})
}
