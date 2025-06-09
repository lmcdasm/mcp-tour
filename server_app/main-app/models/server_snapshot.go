package models

type McpServerInstanceSnapshot struct {
	ID        string           `json:"id"`
	Tools     []ServerTool     `json:"tools,omitempty"`
	Prompts   []ServerPrompt   `json:"prompts,omitempty"`
	Resources []ServerResource `json:"resources,omitempty"`
}

