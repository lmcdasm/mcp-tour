package api

import (
        "net/http"

        "github.com/gin-gonic/gin"
        "mcp-go-server/manager"
        "mcp-go-server/models"
)

// @Summary Add a Tool to MCP Server
// @Tags servers
// @Accept json
// @Produce json
// @Param id path string true "Server ID"
// @Param tool body models.AddToolRequest true "Tool definition payload"
// @Success 200 {object} models.GenericResponse
// @Failure 400 {string} string "invalid tool json"
// @Failure 500 {string} string "failed to add tool"
// @Router /servers/{id}/tools [post]
func AddTools(sm *manager.ServerManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var req models.AddToolRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, "invalid tool json")
			return
		}
		if err := sm.AddTools(id, []models.ServerTool{req.Tool}); err != nil {
			c.String(http.StatusInternalServerError, "failed to add tool")
			return
		}
		c.JSON(http.StatusOK, models.GenericResponse{
			Status:  "success",
			Message: "Tool added to " + id,
		})
	}
}


// GetTools returns the tools for a given server
// @Summary Get Tools for MCP Server
// @Produce json
// @Param id path string true "Server ID"
// @Success 200 {array} models.ServerTool
// @Failure 404 {object} models.GenericResponse
// @Router /servers/{id}/tools [get]
func GetTools(sm *manager.ServerManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		tools, err := sm.GetTools(id)
		if err != nil {
			c.JSON(http.StatusNotFound, models.GenericResponse{
				Status:  "error",
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, tools)
	}
}

