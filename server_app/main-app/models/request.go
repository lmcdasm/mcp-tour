package models


// Request to create a new MCP Server
type CreateMcpServerRequest struct {
	ID      string `json:"id" example:"my-server-1"`
	Addr    string `json:"addr" example:":11000"`
	Version string `json:"version" example:"v1.0.0"`
	Transport string `json:"transport" example:"http"`
	BuildType string `json:"buildtype" example:"binary"`
}

// AddToolRequest represents a single tool being added to a specific MCP Server
type AddToolRequest struct {
	Tool ServerTool `json:"tool"`
}

// AddPromptRequest represents a single prompt being added to a specific MCP Server
type AddPromptRequest struct {
	Prompt ServerPrompt `json:"prompt"`
}

// AddResourceRequest represents a single resource being added to a specific MCP Server
type AddResourceRequest struct {
	Resource ServerResource `json:"resource"`
}

// @Description ServerTool (Swagger-safe)
// NOTE: This avoids jsonschema.Schema or RawMessage parsing errors
type SwaggerServerTool struct {
	Name        string `json:"name" example:"greet"`
	Description string `json:"description" example:"say hi to someone"`
}

// Used ONLY for swagger declaration
type SwaggerAddMcpServerComponentRequest struct {
	ServerID  string                `json:"id"`
	Tools     []SwaggerServerTool  `json:"tools,omitempty"`
	Prompts   []string             `json:"prompts,omitempty"`   // Simplified for docs
	Resources []string             `json:"resources,omitempty"` // Simplified for docs
}
