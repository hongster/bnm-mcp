package main

import (
	"log"

	"github.com/hongster/bnm-mcp/internal/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Start MCP server in STDIO mode.
func main() {
	mcpServer := mcp.NewServer()

	// Start the stdio server
	if err := server.ServeStdio(mcpServer); err != nil {
		log.Fatalf("Server error: %v\n", err)
	}
}
