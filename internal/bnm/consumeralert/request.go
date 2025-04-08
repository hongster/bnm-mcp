package consumeralert

import (
	"github.com/hongster/bnm-mcp/internal/bnm"
)

// Send requesst to https://api.bnm.gov.my/public/consumer-alert.
// Get list of known companies and websites which are neither authorised nor approved under the relevant laws and
// regulations administered by BNM.
func Request(api bnm.Requester) ([]Company, error) {
	var response Response
	err := api.Request("consumer-alert", nil, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
