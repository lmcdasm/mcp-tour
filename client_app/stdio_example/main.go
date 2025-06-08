package main

import (
	"context"
	"log"
	"os/exec"

	"mcp-go-client/mcplib"
)

func main() {
	ctx := context.Background()
	// Create a new client, with no features.
	client := mcp.NewClient("mcp-client", "v1.0.0", nil)
	// Connect to a server over stdin/stdout
	transport := mcp.NewCommandTransport(exec.Command("../../server_app/stdio_example/mcp-go-server"))
	session, err := client.Connect(ctx, transport)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	// Call a tool on the server.
	params := &mcp.CallToolParams[map[string]any]{
		Name:      "greet",
		Arguments: map[string]any{"name": "you"},
	}
	if res, err := mcp.CallTool(ctx, session, params); err != nil {
		log.Printf("CallTool failed: %v", err)
	} else {
		if res.IsError {
			log.Print("tool failed")
		}
		for _, c := range res.Content {
			log.Print(c.Text)
		}
	}
}
