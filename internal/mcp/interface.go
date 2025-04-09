package mcp

import "github.com/mark3labs/mcp-go/server"

// Each MCP tool must provide a Register function to register its tool to the given MCP server.
type RegisterFunc func(*server.MCPServer) error
