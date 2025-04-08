# BNM-MCP

This is a MCP (Model Context Provider) server for BNM (Bank Negara Malaysia) OpenAPI. It only covers a subset of the BNM API.

**DISCLAIMER: This project is neither affiliated with BNM nor endorsed by BNM.**

# MCP Tools

- Financial Consumer Alert
- *More coming soon...*

# Getting Started

STDIO servers are provided for different platforms:
- MacOS: [`build/darwin_arm64/bmn-mcp`](build/darwin_arm64/bnm-mcp)
- Linux: [`build/linux_amd64/bnm-mcp`](build/linux_amd64/bnm-mcp)
- Windows: [`build/windows_amd64/bnm-mcp.exe`](build/windows_amd64/bnm-mcp.exe)

Sample JSON configuration for MCP clients (e.g. Claude Desktop, Cherry Studio):
```json
{
  "mcpServers": {
    "bmn-mcp": {
      "description": "Wrapper for Bank Negara Malaysia OpenAPI.",
      "command": "<PROJECT_PATH>/build/darwin_amd64/bnm-mcp",
      "args": []
    }
  }
}
```

# Development

*TODO*

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.