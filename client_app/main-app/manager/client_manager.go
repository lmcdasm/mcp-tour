package manager

import (
	"context"
	"fmt"
	"sync"

	"mcp-go-client/mcplib"
)

type ClientInstance struct {
	ID       string
	Type     string
	ServerURL string
	Client   *mcp.Client
	Session  *mcp.ClientSession
	Status   string // "running", "stopped"
}

type ClientManager struct {
	mu      sync.Mutex
	clients map[string]*ClientInstance
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		clients: make(map[string]*ClientInstance),
	}
}

func (cm *ClientManager) AddClient(id, clientType, serverURL string) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	if _, exists := cm.clients[id]; exists {
		return fmt.Errorf("client %s already exists", id)
	}

	client := mcp.NewClient(id, "v0.1.0", nil)
	cm.clients[id] = &ClientInstance{
		ID:        id,
		Type:      clientType,
		ServerURL: serverURL,
		Client:    client,
		Status:    "stopped",
	}

	return nil
}

func (cm *ClientManager) StartClient(id string) error {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	inst, ok := cm.clients[id]
	if !ok {
		return fmt.Errorf("client %s not found", id)
	}
	if inst.Status == "running" {
		return nil
	}

	ctx := context.Background()
	transport := mcp.NewSSEClientTransport(inst.ServerURL)
	session, err := inst.Client.Connect(ctx, transport)
	if err != nil {
		return err
	}
	inst.Session = session
	inst.Status = "running"
	return nil
}

func (cm *ClientManager) StopClient(id string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	if inst, ok := cm.clients[id]; ok && inst.Session != nil {
		inst.Session.Close()
		inst.Status = "stopped"
	}
}

func (cm *ClientManager) RemoveClient(id string) {
	cm.StopClient(id)
	delete(cm.clients, id)
}

