package consumeralert

import "testing"

func TestCompany_String(t *testing.T) {
	tests := []struct {
		name     string
		company  Company
		expected string
	}{
		{
			name: "Company with single website",
			company: Company{
				Name:               "Test Company",
				RegistrationNumber: "REG123",
				AddedDate:          "2023-01-01",
				Websites:           []string{"https://example.com"},
			},
			expected: "Name: Test Company\nRegistration Number: REG123\nAdded Date: 2023-01-01\nWebsite: https://example.com",
		},
		{
			name: "Company with multiple websites",
			company: Company{
				Name:               "Multi Web Corp",
				RegistrationNumber: "REG456",
				AddedDate:          "2023-02-01",
				Websites:           []string{"https://example1.com", "https://example2.com"},
			},
			expected: "Name: Multi Web Corp\nRegistration Number: REG456\nAdded Date: 2023-02-01\nWebsite: https://example1.com\nWebsite: https://example2.com",
		},
		{
			name: "Company with no websites",
			company: Company{
				Name:               "No Web Ltd",
				RegistrationNumber: "REG789",
				AddedDate:          "2023-03-01",
				Websites:           []string{},
			},
			expected: "Name: No Web Ltd\nRegistration Number: REG789\nAdded Date: 2023-03-01",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := test.company.String(); got != test.expected {
				t.Errorf("Company.String() = %v, want %v", got, test.expected)
			}
		})
	}
}
