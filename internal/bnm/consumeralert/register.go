package consumeralert

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Register the Financial Consumer Alert tool to the given MCP server.
func Register(mcpServer *server.MCPServer) error {
	// Financial Consumer Alert tool
	mcpTool := mcp.NewTool(
		"consumer_alert",
		mcp.WithDescription("Financial Consumer Alert. Listing of known companies and websites which are neither authorised nor approved under the relevant laws and regulations administered by BNM. (Based on information received by BNM)."),
	)

	// Financial Consumer Alert handler
	mcpServer.AddTool(mcpTool, Handler)

	return nil
}
