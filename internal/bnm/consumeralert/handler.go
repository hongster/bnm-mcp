package consumeralert

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hongster/bnm-mcp/internal/bnm"
	"github.com/mark3labs/mcp-go/mcp"
)

// Execute API request to get list of companies on watch list.
func Handler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	api := bnm.NewAPI(&http.Client{})
	companies, err := Request(api)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("consumeralert handler issue: %s", err)), nil
	}

	return &mcp.CallToolResult{Content: generateCompanyContents(companies)}, nil
}

// Convert company info to text content.
func generateCompanyContents(companies []Company) []mcp.Content {
	var content []mcp.Content
	for _, company := range companies {
		content = append(content, mcp.TextContent{
			Type: "text",
			Text: company.String(),
		})
	}

	return content
}
