package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	ginprom "github.com/zsais/go-gin-prometheus"

	_ "mcp-go-server/docs"
	"mcp-go-server/manager"
)

// @title MCP Explorer - MCP Server APIs
// @version 0.0.3
// @description APIs for MCP Server Instantiation, Configuration and Handling
// @BasePath /

func main() {
	log.Println("MCP Explorer - Server Manager Startup")

	// Create the MCP server manager instance
	serverMgr := manager.NewServerManager()

	// Setup Gin
	mainRouter := gin.Default()
	metricsRouter := gin.Default()

	// Set up Swagger
	mainRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Init API routes
	initializeRoutes(mainRouter, serverMgr)

	// Setup Prometheus metrics
	prom := ginprom.NewPrometheus("mcp")
	prom.Use(metricsRouter)

	// Start metrics server
	go func() {
		log.Println("Metrics server listening on :10011")
		if err := metricsRouter.Run(":10011"); err != nil {
			log.Fatalf("Metrics server error: %v", err)
		}
	}()

	// Start API server
	log.Println("MCP API server listening on :10010")
	if err := mainRouter.Run(":10010"); err != nil {
		log.Fatalf("API server error: %v", err)
	}
}

