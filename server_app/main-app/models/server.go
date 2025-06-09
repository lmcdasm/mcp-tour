package models

type McpServerStatus string

const (
	StatusDefined    McpServerStatus = "defined"
	StatusConfigured McpServerStatus = "configured"
	StatusRunning    McpServerStatus = "running"
	StatusStopped    McpServerStatus = "stopped"
)

type McpServerDefinition struct {
	ID        string `json:"id"`
	Addr      string `json:"addr"`
	Version   string `json:"version"`
	Transport string `json:"transport"` // "StdIO", "Memory", "HTTP"
	BuildType string `json:"build_type"`
}

type McpServerComponents struct {
	Tools     []ServerTool
	Prompts   []ServerPrompt
	Resources []ServerResource
}

type ServerTool struct {
	Name        string
	Description string
	FuncName    string
}

type ServerPrompt struct {
    Name        string
    Description string
    Arguments   []PromptArgument
    HandlerName string
}

type PromptArgument struct {
    Name        string `json:"name"`        // The name of the argument
    Description string `json:"description"` // What this argument is used for
    Type        string `json:"type"`        // Expected type (e.g. "string", "int", "bool")
    Required    bool   `json:"required"`    // Whether this argument is mandatory
}

type ServerResource struct {
    Name        string
    Description string
    URI         string
    MIMEType    string
    Size        int64
}


