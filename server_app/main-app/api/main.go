package api

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"mcp-go-server/manager"
	"mcp-go-server/models"
	//mcplib "mcp-go-server/mcplib"
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
			return
		}
		if err := sm.AddServer(req.ID, req.Addr, req.Version, req.Transport, req.BuildType); err != nil {
			msg := fmt.Sprintf("Failed to Add Server with %v", err)
			c.JSON(http.StatusInternalServerError, models.GenericResponse {
				Status: "internal error",
				Message: msg,
			})
			return
		}

		msg := "MCP Server Instance Created with ID: " + req.ID
		c.JSON(http.StatusCreated, models.GenericResponse {
			Status: "success",
			Message: msg, 
		})
	}
}


// GetMcpServerInstance returns the full in-memory snapshot of a server
// @Summary Get full MCP Server definition
// @Produce json
// @Param id path string true "Server ID"
// @Success 200 {object} models.McpServerInstanceSnapshot
// @Failure 404 {object} models.GenericResponse
// @Router /servers/{id} [get]
func GetMcpServerInstance(sm *manager.ServerManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		snapshot, err := sm.GetServerSnapshot(id)
		if err != nil {
			c.JSON(http.StatusNotFound, models.GenericResponse{
				Status:  "error",
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, snapshot)
	}
}

