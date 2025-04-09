package exchangerate

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Register the Exchange Rate tool to the given MCP server.
func Register(mcpServer *server.MCPServer) error {
	// Exchange Rate tool
	mcpTool := mcp.NewTool(
		"exchange_rate",
		mcp.WithDescription("Latest selling/buying rates for different currencies published by Bank Negara Malaysia. The rates are quoted in Malaysia Ringgit (MYR). This is useful for calculating currency exchange."),
	)

	// Exchange Rate handler
	mcpServer.AddTool(mcpTool, Handler)

	return nil
}
