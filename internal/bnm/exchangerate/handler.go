package exchangerate

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hongster/bnm-mcp/internal/bnm"
	"github.com/mark3labs/mcp-go/mcp"
)

// Execute API request to get list of currencies with corresponding rates.
func Handler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	api := bnm.NewAPI(&http.Client{})
	companies, err := Request(api)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("exchangerate handler issue: %s", err)), nil
	}

	return &mcp.CallToolResult{Content: generateCurrencyContents(companies)}, nil
}

// Convert Currency info to text content.
func generateCurrencyContents(currencies []Currency) []mcp.Content {
	var content []mcp.Content
	for _, currency := range currencies {
		content = append(content, mcp.TextContent{
			Type: "text",
			Text: currency.String(),
		})
	}

	return content
}
