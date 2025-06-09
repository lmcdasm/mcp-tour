package manager

import (
	// Std Libs
	"fmt"
	"os"
	"path/filepath"
	"github.com/google/uuid"

	// Our Stuff
	"mcp-go-server/models"
	"mcp-go-server/logutil"
)

var component_name_factory = "server-manager-factory"
var logr = logutil.InitLogger(component_name_factory)

type MCPServerFactory interface {
	Define(def models.McpServerDefinition) error
	AddComponents(id string, comps models.McpServerComponents) error
	Build(id string) error
	Start(id string) error
	Stop(id string) error
	Status(id string) (models.McpServerStatus, error)
}

type DefaultServerFactory struct {
	id	string // Unique ID of this Factory - since we might want many later
	workspace string
	registry  map[string]*ServerBuildContext
}

func NewDefaultServerFactory(basePath string) *DefaultServerFactory {
	shortID := uuid.New().String()[:7]
	resolvedPath := filepath.Join("/tmp", basePath, shortID)

	// Hard Error - will cause POD / Container to restart if unable to 
	//  create workspace 
	if err := os.MkdirAll(resolvedPath, 0755); err != nil {
		log.Panicf("Failed to Create Workspace Directory: %v", err)
	}

	log.Infof("Factory - Mcp Server Factory created with Id: %v", shortID)

	return &DefaultServerFactory{
		id: shortID,
		workspace: resolvedPath,
		registry:  make(map[string]*ServerBuildContext),
	}

}

func (f *DefaultServerFactory) Define(def models.McpServerDefinition) error {
	logr.Debugf("Define - def.ID: %v", def.ID)
	for k := range f.registry {
    		logr.Debugf("Registry already has key: %v", k)
	}

	if _, exists := f.registry[def.ID]; exists {
		logr.Debugf("f.registry[def.ID] Output: %v", f.registry[def.ID])
		return fmt.Errorf("server ID %s already defined", def.ID)
	}

	ctx := &ServerBuildContext{
		Definition: def,
		Components: models.McpServerComponents{},
		Status:     models.StatusDefined,
	}

	f.registry[def.ID] = ctx
	return nil
}

func (f *DefaultServerFactory) AddComponents(id string, comps models.McpServerComponents) error {
	ctx, ok := f.registry[id]
	if !ok {
		return fmt.Errorf("no such server ID: %s", id)
	}
	ctx.Components.Tools = append(ctx.Components.Tools, comps.Tools...)
	ctx.Components.Prompts = append(ctx.Components.Prompts, comps.Prompts...)
	ctx.Components.Resources = append(ctx.Components.Resources, comps.Resources...)
	ctx.Status = models.StatusConfigured
	return nil
}

func (f *DefaultServerFactory) Build(id string) error {
	ctx, ok := f.registry[id]
	if !ok {
		return fmt.Errorf("server %s not found", id)
	}
	return writeMainGo(ctx, f.workspace)
}

func (f *DefaultServerFactory) Start(id string) error {
	return fmt.Errorf("Start() not yet implemented")
}

func (f *DefaultServerFactory) Stop(id string) error {
	return fmt.Errorf("Stop() not yet implemented")
}

func (f *DefaultServerFactory) Status(id string) (models.McpServerStatus, error) {
	ctx, ok := f.registry[id]
	if !ok {
		return "", fmt.Errorf("server %s not found", id)
	}
	return ctx.Status, nil
}

func (f *DefaultServerFactory) AnnouceToRegistry(serverID, registryURI string) error {
	logr.Infof("Register to MCP Registry (AnnouceToRegistry) not yet implmented - server ID: %v, target URI: %v", serverID, registryURI)
        return nil
}

// Helpers
func (f *DefaultServerFactory) GetDefinition(id string) (*models.McpServerDefinition, error) {
    ctx, ok := f.registry[id]
    if !ok {
        return nil, fmt.Errorf("server %s not found", id)
    }
    return &ctx.Definition, nil
}


func (f *DefaultServerFactory) GetComponents(id string) (*models.McpServerComponents, error) {
    ctx, ok := f.registry[id]
    if !ok {
        return nil, fmt.Errorf("server %s not found", id)
    }
    return &ctx.Components, nil
}

