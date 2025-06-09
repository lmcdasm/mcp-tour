package main

import (
        // 3PP Libs
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
        "github.com/Depado/ginprom"
	"github.com/gin-contrib/cors"



	// Our Stuff
	"mcp-go-server/docs" // Required for Swaggo
	"mcp-go-server/manager"
	"mcp-go-server/logutil"
	"mcp-go-server/sse"
	"mcp-go-server/metrics"
)

// VARS
const version = "0.0.5"
var component_name = "server-main"
var log = logutil.InitLogger(component_name)
var sseManager = sse.NewSSEManager()


// @title MCP Explorer - MCP Server APIs
// @version 0.0.1
// @description APIs for MCP Server Instantiation, Configuration and Handling
// @BasePath /
func main() {
	log.Infof("MCP Explorer - Starting %s, version: %s", component_name, version)
	docs.SwaggerInfo.Version = version

	//Set gin Production release mode
        gin.SetMode(gin.ReleaseMode)

	// Init SSE Manager
	go sseManager.Run()
	log.Infof("MCP Explorer - ServerSideEvent (SSE) Manager started")

	// Init Mcp Server Manager
	serverMgr := manager.NewServerManager("/mcp-workspace")
	log.Infof("MCP Explorer - McpServer Manager started")

	// Primary app router
	mainRouter := gin.Default()

	// Allow CORS 
	mainRouter.Use(cors.Default()) // Allows from all - this can bbe restricted better below

	//mainRouter.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"http://mcp-explorer.svc.local:9000"},
	//	AllowMethods:     []string{"GET", "POST", "OPTIONS"},
	//	AllowHeaders:     []string{"Origin", "Content-Type"},
	//	AllowCredentials: true,
	//	MaxAge:           12 * time.Hour,
	//}))
	// Metrics router (on port 9100)
	metricsRouter := gin.Default()

	// ginprom (prometheus gin wrapper) and expose on MetricsRouter
        p := ginprom.New(
                        ginprom.Engine(metricsRouter),
                        ginprom.Subsystem("gin"),
                        ginprom.Path("/metrics"),
        )

	// Wrap our mainRouter
	mainRouter.Use(p.Instrument())

	// Add our Custom Metrics
	metrics.RegisterCustomMetrics()


	// Add Swagger UI Route
	mainRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ServerManager API Routes
	initializeRoutes(mainRouter, serverMgr)

	// Add SSE Route
	mainRouter.GET("/api/stream", sseManager.Handler())

	// Start Pushing Metrics over SSE Manager
	err := metrics.PushMetricsToSSE(sseManager, serverMgr)
	if err != nil {
		log.Fatalf("failed to start SSE Metrics Push Routing.. : %v", err)
	}

	// Launch metrics server
	go func() {
		log.Infof("Starting metrics server on :9100")
		if err := metricsRouter.Run(":9100"); err != nil {
			log.Fatalf("Metrics server error: %v", err)
		}
	}()

	// Launch main server
	log.Infof("Starting main MCP API on :10010")
	if err := mainRouter.Run(":10010"); err != nil {
		log.Fatalf("Main server error: %v", err)
	}

}

