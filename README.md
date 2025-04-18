# BNM-MCP

This is a MCP (Model Context Provider) server for BNM (Bank Negara Malaysia) OpenAPI. It only covers a subset of the BNM API.

[![Demo video](https://img.youtube.com/vi/7lVRyodWdSI/0.jpg)](https://youtu.be/7lVRyodWdSI?si=WJq1GPv9d-hfg72z)

**DISCLAIMER: This project is neither affiliated with BNM nor endorsed by BNM.**

# MCP Tools

- Financial Consumer Alert
- Exchange Rate
- *more coming soon...*

# Getting Started

STDIO servers are provided for different platforms:
- MacOS: [`build/darwin_amd64/bmn-mcp`](build/darwin_amd64/bnm-mcp)
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

## Prerequisites

- Go 1.24
- Make

## Build & Run

Get the source code
```bash
git clone git@github.com:hongster/bnm-mcp.git
cd bnm-mcp
```

Build the project (for MacOS, Linux, Windows)
```bash 
make build
```

STDIO servers are available in the `build` directory.

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.