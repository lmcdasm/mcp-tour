package metrics

import (
	"mcp-go-server/manager"
)

// McpServerListItem is the structure sent via SSE for each server in the registry
type McpServerListItem struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	BuildType     string `json:"build_type"`
	State         string `json:"state"`
	PromptCount   int    `json:"prompt_count"`
	ToolCount     int    `json:"tool_count"`
	ResourceCount int    `json:"resource_count"`
}

// GetServerListData returns a list of McpServerListItem populated from the factory
func GetServerListData(sm *manager.ServerManager) []McpServerListItem {
	list := []McpServerListItem{}

	for id, ctx := range sm.Factory.Registry {
		def := ctx.Definition
		comps := ctx.Components

		item := McpServerListItem{
			ID:            id,
			Name:          def.ID,
			BuildType:     def.BuildType,
			State:         string(ctx.Status),
			PromptCount:   len(comps.Prompts),
			ToolCount:     len(comps.Tools),
			ResourceCount: len(comps.Resources),
		}

		list = append(list, item)
	}

	return list
}

