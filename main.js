#!/usr/bin/env node

const fs = require('fs');
const path = require('path');
const readline = require('readline');

// Documentation resource cache
const resources = [];

// Load all documentation resources
function loadResources() {
    const docsDir = './docs';

    const walkDir = (dir) => {
        const files = fs.readdirSync(dir, { withFileTypes: true });

        for (const file of files) {
            const filePath = path.join(dir, file.name);

            if (file.isDirectory()) {
                walkDir(filePath);
            } else if (file.name.endsWith('.md')) {
                const content = fs.readFileSync(filePath, 'utf8');
                const name = path.basename(file.name, '.md');
                const { title, description, resourceType, deprecated } = parseDocContent(content);
                const stats = fs.statSync(filePath);

                resources.push({
                    name,
                    title,
                    description,
                    content,
                    type: resourceType,
                    deprecated,
                    last_updated: stats.mtime.toISOString().split('T')[0]
                });
            }
        }
    };

    try {
        walkDir(docsDir);
        resources.sort((a, b) => a.name.localeCompare(b.name));
        console.error(`Loaded ${resources.length} documentation resources`);
    } catch (err) {
        console.error('Error loading resources:', err);
        process.exit(1);
    }
}

function parseDocContent(content) {
    const lines = content.split('\n');
    let title = '';
    let description = '';
    let inDescription = false;
    const descLines = [];

    // Extract title
    for (const line of lines) {
        if (line.startsWith('# ')) {
            title = line.slice(2).trim();
            break;
        }
    }

    // Extract description
    for (const line of lines) {
        if (line.startsWith('## Description')) {
            inDescription = true;
            continue;
        }
        if (inDescription && line.startsWith('##')) break;
        if (inDescription && line.trim()) {
            descLines.push(line.trim());
        }
    }

    description = descLines.join(' ');
    const deprecated = /DEPRECATED/i.test(content);
    const contentUpper = content.toUpperCase();
    const titleUpper = title.toUpperCase();

    let resourceType = 'resource';
    if (/(INTEGRATION)/.test(titleUpper + contentUpper)) resourceType = 'integration';
    else if (/(POLICY)/.test(titleUpper + contentUpper)) resourceType = 'policy';
    else if (/(WEBHOOK)/.test(titleUpper + contentUpper)) resourceType = 'webhook';
    else if (/(WORKER)/.test(titleUpper + contentUpper)) resourceType = 'worker';
    else if (/(STACK)/.test(titleUpper + contentUpper)) resourceType = 'stack';
    else if (/(MODULE)/.test(titleUpper + contentUpper)) resourceType = 'module';
    else if (/(SPACE)/.test(titleUpper + contentUpper)) resourceType = 'space';
    else if (/(USER|ACCESS)/.test(titleUpper + contentUpper)) resourceType = 'access';
    else if (/(SCHEDULED)/.test(titleUpper + contentUpper)) resourceType = 'automation';
    else if (/(CONTEXT)/.test(titleUpper + contentUpper)) resourceType = 'context';

    return { title, description, resourceType, deprecated };
}

function sendResponse(id, result) {
    console.log(JSON.stringify({ jsonrpc: '2.0', id, result }));
}

function sendError(id, code, message, data) {
    console.log(JSON.stringify({
        jsonrpc: '2.0',
        id,
        error: { code, message, data }
    }));
}

function sendToolResult(id, content, isError = false) {
    sendResponse(id, {
        content: [{ type: 'text', text: content }],
        isError
    });
}

function handleInitialize(request) {
    sendResponse(request.id, {
        protocolVersion: '2024-11-05',
        serverInfo: { name: 'spacelift-docs-mcp-server', version: '1.0.0' },
        capabilities: {
            resources: { subscribe: false, listChanged: false },
            tools: { listChanged: false }
        }
    });
}

function handleResourcesList(request) {
    sendResponse(request.id, {
        resources: resources.map(doc => ({
            uri: `spacelift://docs/${doc.name}`,
            name: doc.title,
            description: doc.description,
            mimeType: 'text/markdown'
        }))
    });
}

function handleResourcesRead(request) {
    const { uri } = request.params || {};

    if (!uri?.startsWith('spacelift://docs/')) {
        return sendError(request.id, -32602, 'Invalid params', 'Invalid URI format');
    }

    const resourceName = uri.slice(17);
    const doc = resources.find(r => r.name === resourceName);

    if (!doc) {
        return sendError(request.id, -32602, 'Resource not found', `Resource ${resourceName} not found`);
    }

    sendResponse(request.id, {
        contents: [{ type: 'text', text: doc.content }]
    });
}

function handleToolsList(request) {
    sendResponse(request.id, {
        tools: [
            {
                name: 'search_docs',
                description: 'Search Spacelift documentation for resources matching a query',
                inputSchema: {
                    type: 'object',
                    properties: {
                        query: { type: 'string', description: 'Search query to find relevant documentation' },
                        category: { type: 'string', description: 'Optional category filter (integration, stack, policy, etc.)' },
                        include_deprecated: { type: 'boolean', description: 'Whether to include deprecated resources in results' }
                    },
                    required: ['query']
                }
            },
            {
                name: 'list_categories',
                description: 'List all available resource categories in the Spacelift documentation',
                inputSchema: { type: 'object', properties: {} }
            },
            {
                name: 'get_resource_by_name',
                description: 'Get detailed information about a specific Spacelift resource by name',
                inputSchema: {
                    type: 'object',
                    properties: {
                        name: { type: 'string', description: 'Name of the resource (e.g., \'aws_integration\', \'stack\', \'policy\')' }
                    },
                    required: ['name']
                }
            }
        ]
    });
}

function handleToolsCall(request) {
    const { name, arguments: args } = request.params || {};

    if (!name || !args) {
        return sendError(request.id, -32602, 'Invalid params', 'Expected object with name and arguments');
    }

    switch (name) {
        case 'search_docs':
            return handleSearchDocs(request.id, args);
        case 'list_categories':
            return handleListCategories(request.id, args);
        case 'get_resource_by_name':
            return handleGetResourceByName(request.id, args);
        default:
            return sendError(request.id, -32601, 'Tool not found', `Unknown tool: ${name}`);
    }
}

function handleSearchDocs(id, args) {
    const { query, category, include_deprecated } = args;

    if (!query) {
        return sendError(id, -32602, 'Invalid arguments', 'Missing or invalid query parameter');
    }

    const queryLower = query.toLowerCase();
    const results = resources.filter(r => {
        if (r.deprecated && !include_deprecated) return false;
        if (category && r.type.toLowerCase() !== category.toLowerCase()) return false;

        const searchText = `${r.name} ${r.title} ${r.description} ${r.content}`.toLowerCase();
        return searchText.includes(queryLower);
    });

    let resultText = `Found ${results.length} resources matching '${query}':\n\n`;

    for (const result of results) {
        resultText += `## ${result.title}\n`;
        resultText += `**Type:** ${result.type}\n`;
        if (result.deprecated) resultText += '**Status:** DEPRECATED\n';
        resultText += `**Description:** ${result.description}\n\n`;
    }

    if (!results.length) {
        resultText += 'No resources found matching your query.';
    }

    sendToolResult(id, resultText);
}

function handleListCategories(id, args) {
    const categoryCount = {};

    for (const resource of resources) {
        categoryCount[resource.type] = (categoryCount[resource.type] || 0) + 1;
    }

    let resultText = 'Available resource categories:\n\n';

    for (const category of Object.keys(categoryCount).sort()) {
        resultText += `- **${category}**: ${categoryCount[category]} resources\n`;
    }

    sendToolResult(id, resultText);
}

function handleGetResourceByName(id, args) {
    const { name } = args;

    if (!name) {
        return sendError(id, -32602, 'Invalid arguments', 'Missing or invalid name parameter');
    }

    const resource = resources.find(r => r.name === name);

    if (!resource) {
        return sendToolResult(id, `Resource '${name}' not found`, true);
    }

    let resultText = `# ${resource.title}\n\n`;
    resultText += `**Type:** ${resource.type}\n`;
    if (resource.deprecated) resultText += '**Status:** DEPRECATED\n';
    resultText += `**Last Updated:** ${resource.last_updated}\n\n`;
    resultText += `**Description:** ${resource.description}\n\n`;
    resultText += '**Full Documentation:**\n\n';
    resultText += resource.content;

    sendToolResult(id, resultText);
}

function handleRequest(request) {
    console.error(`Received request: method=${request.method}, params=${JSON.stringify(request.params)}`);

    switch (request.method) {
        case 'initialize':
            return handleInitialize(request);
        case 'resources/list':
            return handleResourcesList(request);
        case 'resources/read':
            return handleResourcesRead(request);
        case 'tools/list':
            return handleToolsList(request);
        case 'tools/call':
            return handleToolsCall(request);
        default:
            return sendError(request.id, -32601, 'Method not found', `Unknown method: ${request.method}`);
    }
}

// Main
loadResources();
console.error('Starting Spacelift Documentation MCP Server');
console.error('Reading JSON-RPC requests from stdin...');

const rl = readline.createInterface({ input: process.stdin });

rl.on('line', (line) => {
    if (!line.trim()) return;

    try {
        const request = JSON.parse(line);
        handleRequest(request);
    } catch (err) {
        sendError(null, -32700, 'Parse error', err.message);
    }
});

rl.on('close', () => {
    console.error('Stdin closed, exiting');
});
