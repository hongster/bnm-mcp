package exchangerate

import (
	"testing"
)

func TestCurrency_String(t *testing.T) {
	tests := []struct {
		name     string
		currency Currency
		want     string
	}{
		{
			name: "normal currency data",
			currency: Currency{
				CurrencyCode: "USD",
				Unit:         1,
				Rate: Rate{
					Date:        "2024-01-01",
					BuyingRate:  4.5678,
					SellingRate: 4.6789,
					MiddleRate:  4.6234,
				},
			},
			want: "Currency Code: USD\n" +
				"Date: 2024-01-01\n" +
				"Unit: 1\n" +
				"Buying Rate: 4.5678\n" +
				"Selling Rate: 4.6789\n" +
				"Middle Rate: 4.6234",
		},
		{
			name: "currency with zero rates",
			currency: Currency{
				CurrencyCode: "EUR",
				Unit:         100,
				Rate: Rate{
					Date:        "2024-01-02",
					BuyingRate:  0,
					SellingRate: 0,
					MiddleRate:  0,
				},
			},
			want: "Currency Code: EUR\n" +
				"Date: 2024-01-02\n" +
				"Unit: 100\n" +
				"Buying Rate: 0\n" +
				"Selling Rate: 0\n" +
				"Middle Rate: 0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.currency.String(); got != tt.want {
				t.Errorf("Currency.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
