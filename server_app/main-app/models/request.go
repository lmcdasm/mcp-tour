package models

import (
	mcplib "mcp-go-server/mcplib"
)

// Request to create a new MCP Server
type CreateMcpServerRequest struct {
	ID      string `json:"id" example:"my-server-1"`
	Addr    string `json:"addr" example:":11000"`
	Version string `json:"version" example:"v1.0.0"`
}

// Request to add components to an existing server
type AddMcpServerComponentRequest struct {
	ServerID  string                `json:"id"`
	Tools     []mcplib.ServerTool         `json:"tools,omitempty"`
	Prompts   []mcplib.ServerPrompt       `json:"prompts,omitempty"`
	Resources []mcplib.ServerResource     `json:"resources,omitempty"`
}

