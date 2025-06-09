package main

import (

	"github.com/gin-gonic/gin"
	"mcp-go-server/manager"
	"mcp-go-server/api"
)


func initializeRoutes(router *gin.Engine, sm *manager.ServerManager) {
    log.Infof("MCP Explorer - Initializing ServerManager API routes..")

    router.GET("/isalive", api.ServerManagerHealthCheck)

    router.POST("/servers/create", api.CreateMcpServerInstance(sm))
    router.GET("/servers/:id", api.GetMcpServerInstance(sm))

    router.POST("/servers/:id/tools", api.AddTools(sm))
    router.POST("/servers/:id/prompts", api.AddPrompts(sm))
    router.POST("/servers/:id/resources", api.AddResources(sm))

    router.GET("/servers/:id/tools", api.GetTools(sm))
    router.GET("/servers/:id/prompts", api.GetPrompts(sm))
    router.GET("/servers/:id/resources", api.GetResources(sm))
}

