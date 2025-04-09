package exchangerate

import (
	"strconv"
	"strings"
)

type Response struct {
	Data []Currency `json:"data"`
	Meta Meta       `json:"meta"`
}

type Currency struct {
	CurrencyCode string `json:"currency_code"`
	Unit         int64  `json:"unit"`
	Rate         Rate   `json:"rate"`
}

type Rate struct {
	Date        string  `json:"date"`
	BuyingRate  float64 `json:"buying_rate"`
	SellingRate float64 `json:"selling_rate"`
	MiddleRate  float64 `json:"middle_rate"`
}

type Meta struct {
	Quote       string `json:"quote"`
	Session     string `json:"session"`
	LastUpdated string `json:"last_updated"`
	TotalResult int    `json:"total_result"`
}

// Format company information as text.
func (c Currency) String() string {
	var builder strings.Builder
	builder.WriteString("Currency Code: " + c.CurrencyCode + "\n")
	builder.WriteString("Date: " + c.Rate.Date + "\n")
	builder.WriteString("Unit: " + strconv.FormatInt(c.Unit, 10) + "\n")
	builder.WriteString("Buying Rate: " + strconv.FormatFloat(c.Rate.BuyingRate, 'f', -1, 64) + "\n")
	builder.WriteString("Selling Rate: " + strconv.FormatFloat(c.Rate.SellingRate, 'f', -1, 64) + "\n")
	builder.WriteString("Middle Rate: " + strconv.FormatFloat(c.Rate.MiddleRate, 'f', -1, 64))
	return builder.String()
}
