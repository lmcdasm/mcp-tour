package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mcp-go-server/manager"
	mcplib "mcp-go-server/mcplib"
)

// Server creation payload
type CreateMcpServerRequest struct {
	ID      string `json:"id" example:"my-server-1"`
	Addr    string `json:"addr" example:":11000"`
	Version string `json:"version" example:"v1.0.0"`
}

// Component addition payload
type AddMcpServerComponentRequest struct {
	ServerID  string                `json:"id"`
	Tools     []mcplib.ServerTool  `json:"tools,omitempty"`
	Prompts   []mcplib.ServerPrompt `json:"prompts,omitempty"`
	Resources []mcplib.ServerResource `json:"resources,omitempty"`
}

func initializeRoutes(router *gin.Engine, sm *manager.ServerManager) {

	// @Summary ServerManager HealthCheck
	// @Tags servers
	// @Success 200 {string} string "OK"
	// @Router /servers [get]
	router.GET("/servers", func(c *gin.Context) {
		c.String(http.StatusOK, "HC - ServerManager Running")
	})

	// @Summary Create MCP Server Instance
	// @Tags servers
	// @Accept json
	// @Produce json
	// @Param request body CreateMcpServerRequest true "Server Definition"
	// @Success 201 {string} string "created"
	// @Failure 400 {string} string "invalid JSON"
	// @Failure 500 {string} string "internal error"
	// @Router /servers/create [post]
	router.POST("/servers/create", func(c *gin.Context) {
		var req CreateMcpServerRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, "invalid JSON")
			return
		}
		if err := sm.AddServer(req.ID, req.Addr, req.Version); err != nil {
			c.String(http.StatusInternalServerError, "failed to add server: %v", err)
			return
		}
		c.String(http.StatusCreated, "server created")
	})

	// @Summary Add Components (tools/prompts/resources) to MCP Server
	// @Tags servers
	// @Accept json
	// @Produce json
	// @Param request body AddMcpServerComponentRequest true "Components"
	// @Success 200 {string} string "components added"
	// @Failure 400 {string} string "invalid JSON"
	// @Failure 500 {string} string "internal error"
	// @Router /servers/add_components [post]
	router.POST("/servers/add_components", func(c *gin.Context) {
		var req AddMcpServerComponentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, "invalid JSON")
			return
		}

		if len(req.Tools) > 0 {
			if err := sm.AddTools(req.ServerID, req.Tools); err != nil {
				c.String(http.StatusInternalServerError, "failed to add tools: %v", err)
				return
			}
		}
		if len(req.Prompts) > 0 {
			if err := sm.AddPrompts(req.ServerID, req.Prompts); err != nil {
				c.String(http.StatusInternalServerError, "failed to add prompts: %v", err)
				return
			}
		}
		if len(req.Resources) > 0 {
			if err := sm.AddResources(req.ServerID, req.Resources); err != nil {
				c.String(http.StatusInternalServerError, "failed to add resources: %v", err)
				return
			}
		}
		c.String(http.StatusOK, "components added")
	})
}

