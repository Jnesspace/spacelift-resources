# Spacelift Documentation MCP Server

A Model Context Protocol (MCP) server that provides GitHub Copilot with access to Spacelift Terraform provider documentation.

## Overview

This MCP server implements the JSON-RPC 2.0 protocol to serve Spacelift documentation to GitHub Copilot and other MCP-compatible clients. It provides searchable access to 42+ Spacelift resources including integrations, stacks, policies, and more.

## Features

- **MCP Protocol Compliance**: Full JSON-RPC 2.0 implementation
- **Rich Documentation Access**: 42+ Spacelift resources with detailed documentation
- **Powerful Search Tools**: Search by query, category, or specific resource names
- **GitHub Copilot Integration**: Seamless integration with VS Code and GitHub Copilot
- **Resource Categories**: Organized by type (integration, stack, policy, webhook, etc.)
- **Deprecation Awareness**: Tracks and filters deprecated resources

## Quick Start

### 1. Start the MCP Server
```bash
go run main.go
```

### 2. Configure VS Code
Add this to your VS Code settings.json:
```json
{
  "github.copilot.chat.experimental.mcp.servers": {
    "spacelift-docs": {
      "command": "go",
      "args": ["run", "main.go"],
      "cwd": "/Users/jakenesler/Documents/GitHub/spacelift-resources"
    }
  }
}
```

### 3. Use with GitHub Copilot
- Ask Copilot about Spacelift resources: "How do I create an AWS integration in Spacelift?"
- Get specific resource info: "Show me the spacelift_stack resource documentation"
- Search documentation: "What are the available Spacelift webhook options?"

## Available Tools

The MCP server provides these tools to GitHub Copilot:

- **search_docs**: Search documentation with optional category filtering
- **list_categories**: Show all available resource categories
- **get_resource_by_name**: Get detailed information about specific resources

## Supported Resource Categories

- **integration**: AWS, Azure, GCP, GitLab, Bitbucket integrations
- **stack**: Stack management and configuration
- **policy**: Access and security policies
- **webhook**: Webhook configurations and secrets
- **worker**: Worker pools and VCS agents
- **module**: Terraform modules
- **space**: Space management
- **context**: Environment contexts and attachments
- **automation**: Scheduled tasks and runs
- **access**: User management and permissions

## Example Usage in VS Code

Once configured, you can ask GitHub Copilot questions like:

```
@spacelift-docs How do I create an AWS integration?
@spacelift-docs What are the required fields for a spacelift_stack?
@spacelift-docs Show me webhook configuration options
@spacelift-docs List all available integration types
```

## Requirements

- Go 1.19+
- VS Code with GitHub Copilot extension
- MCP-compatible client (GitHub Copilot Chat)

## Project Structure

```
.
├── main.go           # MCP server implementation
├── go.mod           # Go module dependencies
├── README.md        # This documentation
└── docs/            # Spacelift documentation files (42 resources)
    ├── aws_integration.md
    ├── stack.md
    ├── policy.md
    └── ... (39 more resources)
```

## Development

The server automatically loads all `.md` files from the `docs/` directory and parses them to extract:
- Resource names and titles
- Descriptions and content
- Resource categories
- Deprecation status
- Last updated timestamps

## Troubleshooting

1. **Server not starting**: Check that Go is installed and `go run main.go` works
2. **Copilot not connecting**: Verify the VS Code settings.json configuration
3. **No documentation found**: Ensure the `docs/` directory contains the markdown files
4. **Search not working**: Try more specific queries or use the available tool names
