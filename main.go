package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// MCP JSON-RPC 2.0 structures
type JSONRPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id,omitempty"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

type JSONRPCResponse struct {
	JSONRPC string        `json:"jsonrpc"`
	ID      interface{}   `json:"id,omitempty"`
	Result  interface{}   `json:"result,omitempty"`
	Error   *JSONRPCError `json:"error,omitempty"`
}

type JSONRPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// MCP Protocol structures
type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type ServerCapabilities struct {
	Resources *ResourceCapabilities `json:"resources,omitempty"`
	Tools     *ToolCapabilities     `json:"tools,omitempty"`
}

type ResourceCapabilities struct {
	Subscribe   bool `json:"subscribe,omitempty"`
	ListChanged bool `json:"listChanged,omitempty"`
}

type ToolCapabilities struct {
	ListChanged bool `json:"listChanged,omitempty"`
}

type InitializeResult struct {
	ProtocolVersion string             `json:"protocolVersion"`
	Capabilities    ServerCapabilities `json:"capabilities"`
	ServerInfo      ServerInfo         `json:"serverInfo"`
}

type Resource struct {
	URI         string      `json:"uri"`
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	MimeType    string      `json:"mimeType,omitempty"`
	Annotations interface{} `json:"annotations,omitempty"`
}

type Tool struct {
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	InputSchema interface{} `json:"inputSchema"`
}

type TextContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type ReadResourceResult struct {
	Contents []TextContent `json:"contents"`
}

type CallToolResult struct {
	Content []TextContent `json:"content"`
	IsError bool          `json:"isError,omitempty"`
}

// Documentation resource structure
type DocResource struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Type        string `json:"type"`
	Deprecated  bool   `json:"deprecated"`
	LastUpdated string `json:"last_updated"`
}

var resources []DocResource

func main() {
	log.SetOutput(os.Stderr) // Send logs to stderr so they don't interfere with JSON-RPC

	// Load all documentation resources
	loadResources()

	log.Printf("Starting Spacelift Documentation MCP Server")
	log.Printf("Loaded %d documentation resources", len(resources))
	log.Printf("Reading JSON-RPC requests from stdin...")

	// Create scanner for stdin
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Parse JSON-RPC request
		var request JSONRPCRequest
		if err := json.Unmarshal([]byte(line), &request); err != nil {
			sendError(nil, -32700, "Parse error", err.Error())
			continue
		}

		// Handle the request
		handleRequest(&request)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from stdin: %v", err)
	}
}

func handleRequest(request *JSONRPCRequest) {
	// Log the method and params of every incoming request
	log.Printf("Received request: method=%s, params=%v", request.Method, request.Params)
	switch request.Method {
	case "initialize":
		handleInitialize(request)
	case "resources/list":
		handleResourcesList(request)
	case "resources/read":
		handleResourcesRead(request)
	case "tools/list":
		handleToolsList(request)
	case "tools/call":
		handleToolsCall(request)
	default:
		sendError(request.ID, -32601, "Method not found", fmt.Sprintf("Unknown method: %s", request.Method))
	}
}

func handleInitialize(request *JSONRPCRequest) {
	result := InitializeResult{
		ProtocolVersion: "2024-11-05",
		ServerInfo: ServerInfo{
			Name:    "spacelift-docs-mcp-server",
			Version: "1.0.0",
		},
		Capabilities: ServerCapabilities{
			Resources: &ResourceCapabilities{
				Subscribe:   false,
				ListChanged: false,
			},
			Tools: &ToolCapabilities{
				ListChanged: false,
			},
		},
	}

	sendResponse(request.ID, result)
}

func handleResourcesList(request *JSONRPCRequest) {
	var mcpResources []Resource

	for _, doc := range resources {
		resource := Resource{
			URI:         fmt.Sprintf("spacelift://docs/%s", doc.Name),
			Name:        doc.Title,
			Description: doc.Description,
			MimeType:    "text/markdown",
		}
		mcpResources = append(mcpResources, resource)
	}

	result := map[string]interface{}{
		"resources": mcpResources,
	}

	sendResponse(request.ID, result)
}

func handleResourcesRead(request *JSONRPCRequest) {
	params, ok := request.Params.(map[string]interface{})
	if !ok {
		sendError(request.ID, -32602, "Invalid params", "Expected object with uri parameter")
		return
	}

	uri, ok := params["uri"].(string)
	if !ok {
		sendError(request.ID, -32602, "Invalid params", "Missing or invalid uri parameter")
		return
	}

	// Extract resource name from URI
	if !strings.HasPrefix(uri, "spacelift://docs/") {
		sendError(request.ID, -32602, "Invalid params", "Invalid URI format")
		return
	}

	resourceName := strings.TrimPrefix(uri, "spacelift://docs/")

	// Find the resource
	for _, doc := range resources {
		if doc.Name == resourceName {
			result := ReadResourceResult{
				Contents: []TextContent{
					{
						Type: "text",
						Text: doc.Content,
					},
				},
			}
			sendResponse(request.ID, result)
			return
		}
	}

	sendError(request.ID, -32602, "Resource not found", fmt.Sprintf("Resource %s not found", resourceName))
}

func handleToolsList(request *JSONRPCRequest) {
	tools := []Tool{
		{
			Name:        "search_docs",
			Description: "Search Spacelift documentation for resources matching a query",
			InputSchema: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"query": map[string]interface{}{
						"type":        "string",
						"description": "Search query to find relevant documentation",
					},
					"category": map[string]interface{}{
						"type":        "string",
						"description": "Optional category filter (integration, stack, policy, etc.)",
					},
					"include_deprecated": map[string]interface{}{
						"type":        "boolean",
						"description": "Whether to include deprecated resources in results",
					},
				},
				"required": []string{"query"},
			},
		},
		{
			Name:        "list_categories",
			Description: "List all available resource categories in the Spacelift documentation",
			InputSchema: map[string]interface{}{
				"type":       "object",
				"properties": map[string]interface{}{},
			},
		},
		{
			Name:        "get_resource_by_name",
			Description: "Get detailed information about a specific Spacelift resource by name",
			InputSchema: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"name": map[string]interface{}{
						"type":        "string",
						"description": "Name of the resource (e.g., 'aws_integration', 'stack', 'policy')",
					},
				},
				"required": []string{"name"},
			},
		},
	}

	result := map[string]interface{}{
		"tools": tools,
	}

	sendResponse(request.ID, result)
}

func handleToolsCall(request *JSONRPCRequest) {
	params, ok := request.Params.(map[string]interface{})
	if !ok {
		sendError(request.ID, -32602, "Invalid params", "Expected object with name and arguments")
		return
	}

	name, ok := params["name"].(string)
	if !ok {
		sendError(request.ID, -32602, "Invalid params", "Missing or invalid name parameter")
		return
	}

	arguments, ok := params["arguments"].(map[string]interface{})
	if !ok {
		sendError(request.ID, -32602, "Invalid params", "Missing or invalid arguments parameter")
		return
	}

	switch name {
	case "search_docs":
		handleSearchDocs(request.ID, arguments)
	case "list_categories":
		handleListCategories(request.ID, arguments)
	case "get_resource_by_name":
		handleGetResourceByName(request.ID, arguments)
	default:
		sendError(request.ID, -32601, "Tool not found", fmt.Sprintf("Unknown tool: %s", name))
	}
}

func handleSearchDocs(id interface{}, args map[string]interface{}) {
	query, ok := args["query"].(string)
	if !ok {
		sendError(id, -32602, "Invalid arguments", "Missing or invalid query parameter")
		return
	}

	category, _ := args["category"].(string)
	includeDeprecated, _ := args["include_deprecated"].(bool)

	var results []DocResource
	queryLower := strings.ToLower(query)

	for _, resource := range resources {
		// Skip deprecated resources unless explicitly included
		if resource.Deprecated && !includeDeprecated {
			continue
		}

		// Filter by category if specified
		if category != "" && strings.ToLower(resource.Type) != strings.ToLower(category) {
			continue
		}

		// Search in name, title, description, and content
		searchText := strings.ToLower(resource.Name + " " + resource.Title + " " + resource.Description + " " + resource.Content)
		if strings.Contains(searchText, queryLower) {
			results = append(results, resource)
		}
	}

	// Format results as text
	var resultText strings.Builder
	resultText.WriteString(fmt.Sprintf("Found %d resources matching '%s':\n\n", len(results), query))

	for _, result := range results {
		resultText.WriteString(fmt.Sprintf("## %s\n", result.Title))
		resultText.WriteString(fmt.Sprintf("**Type:** %s\n", result.Type))
		if result.Deprecated {
			resultText.WriteString("**Status:** DEPRECATED\n")
		}
		resultText.WriteString(fmt.Sprintf("**Description:** %s\n\n", result.Description))
	}

	if len(results) == 0 {
		resultText.WriteString("No resources found matching your query.")
	}

	sendToolResult(id, resultText.String(), false)
}

func handleListCategories(id interface{}, args map[string]interface{}) {
	categoryCount := make(map[string]int)

	for _, resource := range resources {
		categoryCount[resource.Type]++
	}

	var resultText strings.Builder
	resultText.WriteString("Available resource categories:\n\n")

	// Sort categories for consistent output
	var categories []string
	for category := range categoryCount {
		categories = append(categories, category)
	}
	sort.Strings(categories)

	for _, category := range categories {
		count := categoryCount[category]
		resultText.WriteString(fmt.Sprintf("- **%s**: %d resources\n", category, count))
	}

	sendToolResult(id, resultText.String(), false)
}

func handleGetResourceByName(id interface{}, args map[string]interface{}) {
	name, ok := args["name"].(string)
	if !ok {
		sendError(id, -32602, "Invalid arguments", "Missing or invalid name parameter")
		return
	}

	// Find the resource
	for _, resource := range resources {
		if resource.Name == name {
			var resultText strings.Builder
			resultText.WriteString(fmt.Sprintf("# %s\n\n", resource.Title))
			resultText.WriteString(fmt.Sprintf("**Type:** %s\n", resource.Type))
			if resource.Deprecated {
				resultText.WriteString("**Status:** DEPRECATED\n")
			}
			resultText.WriteString(fmt.Sprintf("**Last Updated:** %s\n\n", resource.LastUpdated))
			resultText.WriteString(fmt.Sprintf("**Description:** %s\n\n", resource.Description))
			resultText.WriteString("**Full Documentation:**\n\n")
			resultText.WriteString(resource.Content)

			sendToolResult(id, resultText.String(), false)
			return
		}
	}

	sendToolResult(id, fmt.Sprintf("Resource '%s' not found", name), true)
}

func sendResponse(id interface{}, result interface{}) {
	response := JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      id,
		Result:  result,
	}

	data, _ := json.Marshal(response)
	fmt.Println(string(data))
}

func sendError(id interface{}, code int, message string, data interface{}) {
	response := JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      id,
		Error: &JSONRPCError{
			Code:    code,
			Message: message,
			Data:    data,
		},
	}

	responseData, _ := json.Marshal(response)
	fmt.Println(string(responseData))
}

func sendToolResult(id interface{}, content string, isError bool) {
	result := CallToolResult{
		Content: []TextContent{
			{
				Type: "text",
				Text: content,
			},
		},
		IsError: isError,
	}

	sendResponse(id, result)
}

func loadResources() {
	docsDir := "./docs"

	err := filepath.WalkDir(docsDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && strings.HasSuffix(path, ".md") {
			content, err := os.ReadFile(path)
			if err != nil {
				log.Printf("Error reading file %s: %v", path, err)
				return nil
			}

			name := strings.TrimSuffix(filepath.Base(path), ".md")
			title, description, resourceType, deprecated := parseDocContent(string(content))

			// Get file modification time
			info, _ := d.Info()
			lastUpdated := time.Now().Format("2006-01-02")
			if info != nil {
				lastUpdated = info.ModTime().Format("2006-01-02")
			}

			resource := DocResource{
				Name:        name,
				Title:       title,
				Description: description,
				Content:     string(content),
				Type:        resourceType,
				Deprecated:  deprecated,
				LastUpdated: lastUpdated,
			}

			resources = append(resources, resource)
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Error loading resources: %v", err)
	}

	// Sort resources by name
	sort.Slice(resources, func(i, j int) bool {
		return resources[i].Name < resources[j].Name
	})
}

func parseDocContent(content string) (title, description, resourceType string, deprecated bool) {
	lines := strings.Split(content, "\n")

	// Extract title (first # heading)
	for _, line := range lines {
		if strings.HasPrefix(line, "# ") {
			title = strings.TrimSpace(strings.TrimPrefix(line, "#"))
			break
		}
	}

	// Extract description (content under ## Description)
	inDescription := false
	var descLines []string
	for _, line := range lines {
		if strings.HasPrefix(line, "## Description") {
			inDescription = true
			continue
		}
		if inDescription && strings.HasPrefix(line, "##") {
			break
		}
		if inDescription && strings.TrimSpace(line) != "" {
			descLines = append(descLines, strings.TrimSpace(line))
		}
	}
	description = strings.Join(descLines, " ")

	// Check if deprecated
	deprecated = strings.Contains(strings.ToUpper(content), "DEPRECATED")

	// Determine resource type based on content and name
	contentUpper := strings.ToUpper(content)
	titleUpper := strings.ToUpper(title)

	switch {
	case strings.Contains(titleUpper, "INTEGRATION") || strings.Contains(contentUpper, "INTEGRATION"):
		resourceType = "integration"
	case strings.Contains(titleUpper, "POLICY") || strings.Contains(contentUpper, "POLICY"):
		resourceType = "policy"
	case strings.Contains(titleUpper, "WEBHOOK") || strings.Contains(contentUpper, "WEBHOOK"):
		resourceType = "webhook"
	case strings.Contains(titleUpper, "WORKER") || strings.Contains(contentUpper, "WORKER"):
		resourceType = "worker"
	case strings.Contains(titleUpper, "STACK") || strings.Contains(contentUpper, "STACK"):
		resourceType = "stack"
	case strings.Contains(titleUpper, "MODULE") || strings.Contains(contentUpper, "MODULE"):
		resourceType = "module"
	case strings.Contains(titleUpper, "SPACE") || strings.Contains(contentUpper, "SPACE"):
		resourceType = "space"
	case strings.Contains(titleUpper, "USER") || strings.Contains(contentUpper, "USER") || strings.Contains(titleUpper, "ACCESS"):
		resourceType = "access"
	case strings.Contains(titleUpper, "SCHEDULED") || strings.Contains(contentUpper, "SCHEDULED"):
		resourceType = "automation"
	case strings.Contains(titleUpper, "CONTEXT") || strings.Contains(contentUpper, "CONTEXT"):
		resourceType = "context"
	default:
		resourceType = "resource"
	}

	return
}
