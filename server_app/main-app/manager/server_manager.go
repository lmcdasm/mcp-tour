package manager

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"mcp-go-server/mcplib"
)

type ServerInstance struct {
	ID      string
	Addr    string
	Handler http.Handler
	Server  *mcp.Server
	Status  string // "running" or "stopped"
}

type ServerManager struct {
	mu      sync.Mutex
	servers map[string]*ServerInstance
}

func NewServerManager() *ServerManager {
	return &ServerManager{
		servers: make(map[string]*ServerInstance),
	}
}

func (sm *ServerManager) AddServer(id, addr string, version string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if _, exists := sm.servers[id]; exists {
		return fmt.Errorf("server %s already exists", id)
	}

	srv := mcp.NewServer(id, version, nil)

	handler := mcp.NewSSEHandler(func(r *http.Request) *mcp.Server {
		return srv
	})

	sm.servers[id] = &ServerInstance{
		ID:      id,
		Addr:    addr,
		Handler: handler,
		Server:  srv,
		Status:  "stopped",
	}

	return nil
}

func (sm *ServerManager) StartServer(id string) error {
	sm.mu.Lock()
	inst, exists := sm.servers[id]
	sm.mu.Unlock()
	if !exists {
		return fmt.Errorf("server %s not found", id)
	}
	if inst.Status == "running" {
		return nil
	}

	go func() {
		log.Printf("Starting MCP Server [%s] on %s", id, inst.Addr)
		err := http.ListenAndServe(inst.Addr, inst.Handler)
		if err != nil {
			log.Printf("Server %s failed: %v", id, err)
		}
	}()
	inst.Status = "running"
	return nil
}

func (sm *ServerManager) RemoveServer(id string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	// Currently cannot forcibly shutdown http.ListenAndServe without embedding
	delete(sm.servers, id)
}


//  Helper Functions to Add Prompts/Resources and Tools to a NewServer that has been defined.
func (sm *ServerManager) AddTools(id string, tools []mcp.ServerTool) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	srv, ok := sm.servers[id]
	if !ok {
		return fmt.Errorf("server %s not found", id)
	}
	toolPtrs := make([]*mcp.ServerTool, len(tools))
	for i := range tools {
		toolPtrs[i] = &tools[i]
	}
	srv.Server.AddTools(toolPtrs...)
	return nil
}

func (sm *ServerManager) AddPrompts(id string, prompts []mcp.ServerPrompt) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	srv, ok := sm.servers[id]
	if !ok {
		return fmt.Errorf("server %s not found", id)
	}
	promptPtrs := make([]*mcp.ServerPrompt, len(prompts))
	for i := range prompts {
		promptPtrs[i] = &prompts[i]
	}
	srv.Server.AddPrompts(promptPtrs...)
	return nil
}

func (sm *ServerManager) AddResources(id string, resources []mcp.ServerResource) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	srv, ok := sm.servers[id]
	if !ok {
		return fmt.Errorf("server %s not found", id)
	}
	resourcePtrs := make([]*mcp.ServerResource, len(resources))
	for i := range resources {
		resourcePtrs[i] = &resources[i]
	}
	srv.Server.AddResources(resourcePtrs...)
	return nil
}
