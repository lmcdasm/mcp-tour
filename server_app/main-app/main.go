package main

// @title MCP Explorer - MCP Server APIs
// @version 0.0.2
// @description APIs for MCP Server Instantiation, Configuration and Handling
// @BasePath / 

import (
	// Std libs
	"log"
	"net/http"
        "context"

	// 3PP libs
	httpSwagger "github.com/swaggo/http-swagger"

	// Our Libs
	_ "mcp-go-server/docs" // Docs is geneated by SWAG CLI during build
        "mcp-go-server/mcplib"
        "mcp-go-server/manager"

 	
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
	log.Printf("MCP Explorer - Server Manager Startup")

        // Create a MCP Server Manager
        serverMgr := manager.NewServerManager()
	log.Printf("MAIN - Creating new ServerManager")


	// Setup our Routes
	initializeRoutes(serverMgr)

	// Server Swagger UI at /swagger/index.html
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)


	// Wrap this in a SSE handler
	//handler := mcp.NewSSEHandler(func(r *http.Request) *mcp.Server {
	//	return server
	//})

	// Start HTTP server up
	addr := ":10010"
	log.Printf("MAIN - MCP SSE server listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
