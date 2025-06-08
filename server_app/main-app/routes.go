package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"mcp-go-server/manager"
	"mcp-go-server/mcplib"
)

///// DASM - Add your Routes here, and then your named Function and Swaggo below
func initializeRoutes(sm *manager.ServerManager) {
    http.HandleFunc("/servers", ServerManagerHealthCheck(sm)) 
    http.HandleFunc("/servers/create", CreateMcpServerInstance(sm))
    http.HandleFunc("/servers/add_components", AddMcpServerComponents(sm))
}



//// DASM - Types (structs) for request information types 

// CreateMcpServerRequest defines the expected body for /servers/create
type CreateMcpServerRequest struct {
	ID     string `json:"id"`
	Addr   string `json:"addr"`
	Version string `json:"version"`
}

// AddMcpServerComponentRequest defines the expected body for /servers/add_components
type AddMcpServerComponentRequest struct {
	ServerID 	string `json:"id"`
	Tools 	[]mcp.ServerTool  `json:"tools,omitempty"`
        Prompts  []mcp.ServerPrompt `json:"prompts,omitempty"`
       Resources []mcp.ServerResource `json:"resources,omitempty"`
} 




//// DASM - Add the matching Handler Fucntions for your routes below here.

// ServerManagerHealthCheck  godoc
// @Summary ServerManager HealthCheck
// @Description Basic Health Check to verify ServerManager is running
// @Tags servers
// @Success 200 {string} string "HC - ServerManager Running"
// @Router /servers [get]
func ServerManagerHealthCheck(sm *manager.ServerManager)  http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HC - ServerManager Running\n"))
	}
}




// CreateMcpServerInstance godoc
// @Summary Create (define)  a new MCP Server Instance to the System
// @Description Creates (Defines) a blank MCP Server Instance with a Unique ID and Address in Runtime
// @Tags servers
// @Accept json
// @Produce json
// @Param request body CreateMcpServerRequest true "Create McpServer Request Input"
// @Success 201 {string} string "MCP Server - Create Server Instance - Success - Instance Created"
// @Failure 400 {string} string "MCP Server - Create Server Instance - Failed - Invalid JSON"
// @Failure 500 {string} string "MCP Server - Create Server Instance - Failed - Error Adding Server"
// @Router /servers/create [post]
func CreateMcpServerInstance(sm *manager.ServerManager)  http.HandlerFunc { 
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req struct {
			ID     string `json:"id"`
			Addr   string `json:"addr"`
			Version string `json:"version"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		if err := sm.AddServer(req.ID, req.Addr, req.Version); err != nil {
			http.Error(w, fmt.Sprintf("failed to add server: %v", err), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("server created\n"))
	}
}

// AddMcpServerComponents godoc
// @Summary Add MCP Component to Server (Prompt, Resource or Tool)
// @Description Adds Prompt, Resource and Tool elements to a given MCP Server Instance
// @Tags servers
// @Accept json
// @Produce json
// @Param request body AddMcpServerComponentRequest true "Add McpServer Component Request Input"
// @Success 201 {string} string "MCP Server - Add Server Components - Successfully added to Instance"
// @Failure 400 {string} string "MCP Server - Add Server Components Failed - Invalid JSON"
// @Failure 500 {string} string "MCP Server - Add Server Components - Failed to Add Components to Server"
// @Router /servers/add_components [post]
func AddMcpServerComponents(sm *manager.ServerManager) http.HandlerFunc {
	return  func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req struct {
			ServerID  string                   `json:"id"`
			Tools     []mcp.ServerTool         `json:"tools,omitempty"`
			Prompts   []mcp.ServerPrompt       `json:"prompts,omitempty"`
			Resources []mcp.ServerResource     `json:"resources,omitempty"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		if len(req.Tools) > 0 {
			if err := sm.AddTools(req.ServerID, req.Tools); err != nil {
				http.Error(w, fmt.Sprintf("failed to add tools: %v", err), http.StatusInternalServerError)
				return
			}
		}

		if len(req.Prompts) > 0 {
			if err := sm.AddPrompts(req.ServerID, req.Prompts); err != nil {
				http.Error(w, fmt.Sprintf("failed to add prompts: %v", err), http.StatusInternalServerError)
				return
			}
		}

		if len(req.Resources) > 0 {
			if err := sm.AddResources(req.ServerID, req.Resources); err != nil {
				http.Error(w, fmt.Sprintf("failed to add resources: %v", err), http.StatusInternalServerError)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("components added\n"))
	}
}


