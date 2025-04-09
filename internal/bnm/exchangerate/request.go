package exchangerate

import "github.com/hongster/bnm-mcp/internal/bnm"

// Send request to https://api.bnm.gov.my/public/exchange-rate.
// Get list of currencies with corresponding rates.
func Request(api bnm.Requester) ([]Currency, error) {
	var response Response
	err := api.Request("exchange-rate", nil, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
