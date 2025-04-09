package mcp

import (
	"log"

	"github.com/hongster/bnm-mcp/internal/bnm/consumeralert"
	"github.com/hongster/bnm-mcp/internal/bnm/exchangerate"
	"github.com/mark3labs/mcp-go/server"
)

// Create a new MCP server. It is not running yet, need `SSEServer` or `STDIOServer` to serve it.
func NewServer() *server.MCPServer {
	// TODO start server in SSE and STDIO mode
	mcpServer := server.NewMCPServer(
		"Bank Negara Malaysia MCP",
		"0.2.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	// Register tools
	registerFuncs := []RegisterFunc{
		consumeralert.Register,
		exchangerate.Register,
	}
	for _, registerFunc := range registerFuncs {
		err := registerFunc(mcpServer)
		if err != nil {
			log.Printf("register tool issue: %s", err)
		}
	}

	return mcpServer
}
