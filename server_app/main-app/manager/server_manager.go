package manager

import (
	"fmt"
	"net/http"
	"sync"

	"mcp-go-server/mcplib"
	"mcp-go-server/models"
	"mcp-go-server/logutil"
)

var component_name_main = "server-manager-main"
var log = logutil.InitLogger(component_name_main)

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

	factory *DefaultServerFactory 
}

func NewServerManager(workspace string) *ServerManager {
	return &ServerManager{
		servers: make(map[string]*ServerInstance),
		factory: NewDefaultServerFactory(workspace),
	}
}

func (sm *ServerManager) AddServer(id, addr, version, transport, buildtype string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// Create a definition using the full model structure
	def := models.McpServerDefinition{
		ID:        id,
		Addr:      addr,
		Version:   version,
		Transport: transport,       // for now hardcode, later accept from CreateRequest
		BuildType: buildtype,    // optional hook
	}

	if err := sm.factory.Define(def); err != nil {
		return fmt.Errorf("factory define failed : %v", err)
	}

	// ServerManager has the lightweight reference of the McpServerInstance (without going deep) to the factory
	sm.servers[id] = &ServerInstance{
		ID:     id,
		Addr:   addr,
		Status: "defined",
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

func (sm *ServerManager) AddTools(id string, tools []models.ServerTool) error {
    ctx, ok := sm.factory.registry[id]
    if !ok {
        return fmt.Errorf("server ID %s not found", id)
    }

    // Initialize slice if needed
    if ctx.Components.Tools == nil {
        ctx.Components.Tools = []models.ServerTool{}
    }

    for _, t := range tools {
        // Add metadata-only tool (defer full tool materialization to WriteMainGo)
        ctx.Components.Tools = append(ctx.Components.Tools, models.ServerTool{
            Name:        t.Name,
            Description: t.Description,
            FuncName:    fmt.Sprintf("%sFunc", t.Name), // convention
        })
    }

    if ctx.Definition.BuildType == "binary" {
        return writeMainGo(ctx, sm.factory.workspace)
    }

    return nil
}
func (sm *ServerManager) AddPrompts(id string, prompts []models.ServerPrompt) error {
    ctx, ok := sm.factory.registry[id]
    if !ok {
        return fmt.Errorf("server ID %s not found", id)
    }

    if ctx.Components.Prompts == nil {
        ctx.Components.Prompts = []models.ServerPrompt{}
    }

    for _, p := range prompts {
        ctx.Components.Prompts = append(ctx.Components.Prompts, models.ServerPrompt{
            Name:        p.Name,
            Description: p.Description,
            Arguments:   p.Arguments,
            HandlerName: fmt.Sprintf("%sPromptHandler", p.Name), // convention
        })
    }

    if ctx.Definition.BuildType == "binary" {
        return writeMainGo(ctx, sm.factory.workspace)
    }

    return nil
}

func (sm *ServerManager) AddResources(id string, resources []models.ServerResource) error {
    ctx, ok := sm.factory.registry[id]
    if !ok {
        return fmt.Errorf("server ID %s not found", id)
    }

    if ctx.Components.Resources == nil {
        ctx.Components.Resources = []models.ServerResource{}
    }

    for _, r := range resources {
        ctx.Components.Resources = append(ctx.Components.Resources, models.ServerResource{
            Name:        r.Name,
            Description: r.Description,
            URI:         r.URI,
            MIMEType:    r.MIMEType,
            Size:        r.Size,
        })
    }

    if ctx.Definition.BuildType == "binary" {
        return writeMainGo(ctx, sm.factory.workspace)
    }

    return nil
}


func (sm *ServerManager) GetTools(id string) ([]models.ServerTool, error) {
	ctx, ok := sm.factory.registry[id]
	if !ok {
		return nil, fmt.Errorf("server ID %s not found", id)
	}
	return ctx.Components.Tools, nil
}

func (sm *ServerManager) GetPrompts(id string) ([]models.ServerPrompt, error) {
	ctx, ok := sm.factory.registry[id]
	if !ok {
		return nil, fmt.Errorf("server ID %s not found", id)
	}
	return ctx.Components.Prompts, nil
}

func (sm *ServerManager) GetResources(id string) ([]models.ServerResource, error) {
	ctx, ok := sm.factory.registry[id]
	if !ok {
		return nil, fmt.Errorf("server ID %s not found", id)
	}
	return ctx.Components.Resources, nil
}


func (sm *ServerManager) GetServerSnapshot(id string) (*models.McpServerInstanceSnapshot, error) {
	ctx, ok := sm.factory.registry[id]
	if !ok {
		return nil, fmt.Errorf("server ID %s not found", id)
	}
	return &models.McpServerInstanceSnapshot{
		ID:        id,
		Tools:     ctx.Components.Tools,
		Prompts:   ctx.Components.Prompts,
		Resources: ctx.Components.Resources,
	}, nil
}

