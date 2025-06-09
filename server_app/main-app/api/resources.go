package api 

import (
        "net/http"

        "github.com/gin-gonic/gin"
        "mcp-go-server/manager"
        "mcp-go-server/models"
)


// @Summary Add a Resource to MCP Server
// @Tags servers
// @Accept json
// @Produce json
// @Param id path string true "Server ID"
// @Param resource body models.AddResourceRequest true "Resource definition payload"
// @Success 200 {object} models.GenericResponse
// @Failure 400 {string} string "invalid resource json"
// @Failure 500 {string} string "failed to add resource"
// @Router /servers/{id}/resources [post]
func AddResources(sm *manager.ServerManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var req models.AddResourceRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.String(http.StatusBadRequest, "invalid resource json")
			return
		}
		if err := sm.AddResources(id, []models.ServerResource{req.Resource}); err != nil {
			c.String(http.StatusInternalServerError, "failed to add resource")
			return
		}
		c.JSON(http.StatusOK, models.GenericResponse{
			Status:  "success",
			Message: "Resource added to " + id,
		})
	}
}

// GetResources returns the resources for a given server
// @Summary Get Resources for MCP Server
// @Produce json
// @Success 200 {array} models.ServerResource
// @Param id path string true "Server ID"
// @Failure 404 {object} models.GenericResponse
// @Router /servers/{id}/resources [get]
func GetResources(sm *manager.ServerManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		resources, err := sm.GetResources(id)
		if err != nil {
			c.JSON(http.StatusNotFound, models.GenericResponse{
				Status:  "error",
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, resources)
	}
}

