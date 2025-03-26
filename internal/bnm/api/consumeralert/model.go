// JSON response data struct
package consumeralert

import "strings"

type Response struct {
	Data []Company `json:"data"`
	Meta Meta      `json:"meta"`
}

type Company struct {
	Name               string   `json:"name"`
	RegistrationNumber string   `json:"regisration_number"`
	AddedDate          string   `json:"added_date"`
	Websites           []string `json:"websites"`
}

type Meta struct {
	LastUpdated string `json:"last_updated"`
	TotalResult int    `json:"total_result"`
}

// For company information as text.
func (c Company) String() string {
	var builder strings.Builder
	builder.WriteString("Name: " + c.Name + "\n")
	builder.WriteString("Registration Number: " + c.RegistrationNumber + "\n")
	builder.WriteString("Added Date: " + c.AddedDate)

	for _, website := range c.Websites {
		builder.WriteString("\nWebsite: " + website)
	}

	return builder.String()
}
