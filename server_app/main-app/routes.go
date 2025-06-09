package main

import (

	"github.com/gin-gonic/gin"
	"mcp-go-server/manager"
	"mcp-go-server/api"
)


func initializeRoutes(router *gin.Engine, sm *manager.ServerManager) {
    log.Infof("MCP Explorer - Initializing ServerManager API routes..")
    router.GET("/servers", api.ServerManagerHealthCheck)
    router.POST("/servers/create", api.CreateMcpServerInstance(sm))
    router.POST("/servers/add_components", api.AddMcpServerComponents(sm))
}

