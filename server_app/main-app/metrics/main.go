package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"

	// Our Stuff
	sse "mcp-go-server/sse"
	"mcp-go-server/logutil"
	"mcp-go-server/manager"
)

var component_name = "server-metrics-main"
var log = logutil.InitLogger(component_name)

// RegisterCustomMetrics registers all app-level Prometheus metrics
func RegisterCustomMetrics() {
	prometheus.MustRegister(
		McpServerGauge,
		McpServerState,
		McpClientGauge,
		McpClientState,
		McpRegistryGauge,
		McpRegistryState,
	)

}

// PushMetricsToSSE sends periodic SSE updates to connected clients
func PushMetricsToSSE(s *sse.SSEManager, sm *manager.ServerManager) error {
	go func() {
		for {
			log.Debugf("pushMetrics: Sleeping ....")
			time.Sleep(5 * time.Second)
			log.Debugf("pushMetrics: Pushing latest Metrics ")

			// Send simple Metric (will remove)
			s.Broadcast(sse.SSEMessage{
				MsgID:   "gauge_server_ops",
				Payload: McpServerGaugeValue(),
			})

			// Broadcast List Metrics
			s.Broadcast(sse.SSEMessage{
				MsgID: "list_server_ops",
				Payload: GetServerListData(sm),
			})
		}
	}()
	return nil
}

