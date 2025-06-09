package main

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	"github.com/zsais/go-gin-prometheus"

	_ "mcp-go-server/docs" // Required for Swaggo

	"mcp-go-server/manager"
	"mcp-go-server/logutil"
)

const version = "0.0.3"
var component_name = "server-manager-main"
var log = logutil.InitLogger(component_name)

// @title MCP Explorer - MCP Server APIs
// @version 0.0.3
// @description APIs for MCP Server Instantiation, Configuration and Handling
// @BasePath /
func main() {
	log.Infof("MCP Explorer - Starting %s, version: %s", component_name, version)

	// Init manager
	serverMgr := manager.NewServerManager()

	// Primary app router
	mainRouter := gin.Default()
	// Metrics router (on port 9100)
	metricsRouter := gin.Default()

	// Attach Prometheus metrics to metrics router
	prom := ginprometheus.NewPrometheus("mcp_metrics")
	prom.Use(metricsRouter)

	// Swagger UI
	mainRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Routes
	initializeRoutes(mainRouter, serverMgr)

	// Launch main server
	go func() {
		log.Infof("Starting main MCP API on :10010")
		if err := mainRouter.Run(":10010"); err != nil {
			log.Fatalf("Main server error: %v", err)
		}
	}()

	// Launch metrics server
	log.Infof("Starting metrics server on :9100")
	if err := metricsRouter.Run(":9100"); err != nil {
		log.Fatalf("Metrics server error: %v", err)
	}
}

