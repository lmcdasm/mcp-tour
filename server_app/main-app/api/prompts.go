package api


import (
        "net/http"

        "github.com/gin-gonic/gin"
        "mcp-go-server/manager"
        "mcp-go-server/models"
)

// @Summary Add a Prompt to MCP Server
// @Tags servers
// @Accept json
// @Produce json
// @Param id path string true "Server ID"
// @Param prompt body models.AddPromptRequest true "Prompt definition payload"
// @Success 200 {object} models.GenericResponse
// @Failure 400 {string} string "invalid prompt json"
// @Failure 500 {string} string "failed to add prompt"
// @Router /servers/{id}/prompts [post]
func AddPrompts(sm *manager.ServerManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var req models.AddPromptRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			msg := "Failed to Add Prompt - Invalid JSON Input for Server ID:  " + id
                	c.JSON(http.StatusBadRequest, models.GenericResponse {
                       	 	Status: "failed",
                        	Message: msg,
                	})
			return
		}
		if err := sm.AddPrompts(id, []models.ServerPrompt{req.Prompt}); err != nil {
			msg := "Failed to Add Prompt to Server: " + id
                	c.JSON(http.StatusInternalServerError, models.GenericResponse {
                       	 	Status: "success",
                        	Message: msg,
                	})
			return
		}
		msg := "Added Prompt to MCP Server ID: " + id 
               	c.JSON(http.StatusOK, models.GenericResponse {
                	Status: "success",
                      	Message: msg,
               	})
	}
}

// GetPrompts returns the prompts for a given server
// @Summary Get Prompts for MCP Server
// @Produce json
// @Param id path string true "Server ID"
// @Success 200 {array} models.ServerPrompt
// @Failure 404 {object} models.GenericResponse
// @Router /servers/{id}/prompts [get]
func GetPrompts(sm *manager.ServerManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		prompts, err := sm.GetPrompts(id)
		if err != nil {
			c.JSON(http.StatusNotFound, models.GenericResponse{
				Status:  "error",
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, prompts)
	}
}

