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
        client := mcp.NewClient("mcp-go-client", "v0.0.2", nil)
        // Connect to a server over sse/http
        transport := mcp.NewCommandTransport(exec.Command("../../server_app/stdio_example/mcp-go-server"))
	transport := mcp.NewSSEClientTransport("http://192.168.1.131:10010/")
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

