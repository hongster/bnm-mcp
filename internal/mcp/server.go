package mcp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hongster/bnm-mcp/internal/bnm"
	"github.com/hongster/bnm-mcp/internal/bnm/consumeralert"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Create a new MCP server. It is not running yet, need `SSEServer` or `STDIOServer` to serve it.
func NewServer() *server.MCPServer {
	// TODO start server in SSE and STDIO mode
	mcpServer := server.NewMCPServer(
		"Bank Negara Malaysia MCP",
		"0.1.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	// Financial Consumer Alert tool
	consumerAlert := mcp.NewTool(
		"consumer_alert",
		mcp.WithDescription("Financial Consumer Alert. Listing of known companies and websites which are neither authorised nor approved under the relevant laws and regulations administered by BNM. (Based on information received by BNM)."),
	)

	// Financial Consumer Alert handler
	mcpServer.AddTool(consumerAlert, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		api := bnm.NewAPI(&http.Client{})
		companies, err := consumeralert.Request(api)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("unable to query BNM API: %s", err)), nil
		}

		return &mcp.CallToolResult{Content: generateCompanyContents(companies)}, nil
	})

	return mcpServer
}

func generateCompanyContents(companies []consumeralert.Company) []mcp.Content {
	var content []mcp.Content
	for _, company := range companies {
		content = append(content, mcp.TextContent{
			Type: "text",
			Text: company.String(),
		})
	}

	return content
}
