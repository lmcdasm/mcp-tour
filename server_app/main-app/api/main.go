package api

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"mcp-go-server/manager"
	"mcp-go-server/models"
)

// @Summary ServerManager HealthCheck
// @Tags servers
// @Success 200 {object} models.HcResponse "alive - HC Passed"
// @Router /servers [get]
func ServerManagerHealthCheck(c *gin.Context) {
	resp := models.HcResponse{
		Status: "alive",
		Message: "HealthCheck Passed",
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary Create MCP Server Instance
// @Tags servers
// @Accept json
// @Produce json
// @Param request body models.CreateMcpServerRequest true "Server Definition"
// @Success 201 {object} models.GenericResponse "created"
// @Failure 400 {object} models.GenericResponse "invalid JSON"
// @Failure 500 {object} models.GenericResponse "internal error"
// @Router /servers/create [post]
func CreateMcpServerInstance(sm *manager.ServerManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.CreateMcpServerRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, models.GenericResponse {
				Status: "failed",
				Message: "invalid JSON input", 
			})
		}
		if err := sm.AddServer(req.ID, req.Addr, req.Version); err != nil {
			msg := fmt.Sprintf("Failed to Add Server with %v", err)
			c.JSON(http.StatusInternalServerError, models.GenericResponse {
				Status: "internal error",
				Message: msg,
			})
		}

		msg := "MCP Server Instance Created with ID: " + req.ID
		c.JSON(http.StatusCreated, models.GenericResponse {
			Status: "success",
			Message: msg, 
		})
	}
}

// @Summary Add Components (tools/prompts/resources) to MCP Server
// @Tags servers
// @Accept json
// @Produce json
// @Param request body models.AddMcpServerComponentRequest true "Components"
// @Success 200 {string} models.GenericResponse "components added"
// @Failure 400 {string} models.GenericResponse "invalid JSON"
// @Failure 500 {string} models.GenericResponse "internal error"
// @Router /servers/add_components [post]
func AddMcpServerComponents(sm *manager.ServerManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.AddMcpServerComponentRequest
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
		msg := "MCP Server Instance " + req.ServerID + " Components Updated"
		c.JSON(http.StatusOK, models.GenericResponse {
			Status: "success",
			Message: msg, 
		})
	}
}

