package main

import (
	"context"

	"mcp-go-server/mcplib"
)

type HiParams struct {
	Name string `json:"name"`
}

func SayHi(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParams[HiParams]) (*mcp.CallToolResult, error) {
	return &mcp.CallToolResult{
		Content: []*mcp.Content{mcp.NewTextContent("Hi " + params.Name)},
	}, nil
}

func main() {
	// Create a server with a single tool.
	server := mcp.NewServer("greeter", "v1.0.0", nil)
	server.AddTools(mcp.NewTool("greet", "say hi", SayHi))
	// Run the server over stdin/stdout, until the client disconnects
	_ = server.Run(context.Background(), mcp.NewStdIOTransport())
}
