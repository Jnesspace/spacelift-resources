# Spacelift Documentation MCP Server

Super quick unofficial reference of terraform resources to save time working with AI agents, saving time looking at the resources when configuring spacelift with Tofu/Terraform

Fork and add to your local machine. 

Tell agent to send requests using the terminal in agent mode, with any LLM IDE

If you cant get it to work ask the agent how to use it :) 
---

## 🚀 Quick Start

### 1. Run the MCP Server and Test from Terminal


```sh
echo '{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"list_categories","arguments":{}}}' | go run main.go
```

### 2. Example Prompts for Agents
- Ask Copilot: `How do I create an AWS integration in Spacelift?`
- Get resource info: `Show me the spacelift_stack resource documentation`
- Search docs: `What are the available Spacelift webhook options?`

---

## 🛠️ Features
- **MCP Protocol Compliance:** Full JSON-RPC 2.0 implementation
- **Rich Documentation Access:** 40+ Spacelift resources with detailed docs
- **Powerful Search Tools:** Query by keyword, category, or resource name
- **Agent Integration:** Works with Copilot, Cursor, LibreChat, and more
- **Resource Categories:** Organized by type (integration, stack, policy, etc.)
- **Deprecation Awareness:** Tracks and filters deprecated resources

---

## 🤖 Available Tools

| Tool Name             | Description                                                      | Example Prompt                                 |
|----------------------|------------------------------------------------------------------|------------------------------------------------|
| `search_docs`        | Search docs for resources matching a query                       | `@spacelift-docs search_docs query="aws"`    |
| `list_categories`    | List all available resource categories                           | `@spacelift-docs list_categories`              |
| `get_resource_by_name`| Get detailed info about a specific resource by name              | `@spacelift-docs get_resource_by_name name="stack"` |

---

## 🗂️ Supported Resource Categories
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

---


## 📝 Project Structure

```
.
├── main.go           # MCP server implementation
├── go.mod            # Go module dependencies
├── README.md         # This documentation
└── docs/             # Spacelift documentation files (markdown)
```

---

## 🧑‍💻 Development Notes
- All `.md` files in `docs/` are loaded as resources.
- Tool schemas and logic are defined in `main.go`.
- To add or change tools, edit the `handleToolsList` and related handler functions.

