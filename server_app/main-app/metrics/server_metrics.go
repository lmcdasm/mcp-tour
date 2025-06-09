package metrics

import (

	"github.com/prometheus/client_golang/prometheus"
	"mcp-go-server/manager"
	"mcp-go-server/models"
)

// --- Gauges ---

var (
	// Total number of defined MCP servers
	McpServerGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mcp",
			Subsystem: "server",
			Name:      "server_ops_total",
			Help:      "Total number of defined MCP servers",
		},
	)

	// Number of servers per lifecycle state (defined, configured, running, stopped)
	McpServerState = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "mcp",
			Subsystem: "server",
			Name:      "server_state_count",
			Help:      "Current count of MCP servers in each state",
		},
		[]string{"status"},
	)

	serverGaugeVal float64
)

// Increment total server counter
func IncMcpServerGauge() {
	serverGaugeVal++
	McpServerGauge.Set(serverGaugeVal)
}

// Decrement total server counter
func DecMcpServerGauge() {
	if serverGaugeVal > 0 {
		serverGaugeVal--
	}
	McpServerGauge.Set(serverGaugeVal)
}

// Return current gauge value
func McpServerGaugeValue() float64 {
	return serverGaugeVal
}

// UpdateServerMetrics pulls from the factory and refreshes the metrics
func UpdateServerMetrics(factory *manager.DefaultServerFactory) {
	statusCounts := map[string]float64{
		string(models.StatusDefined):    0,
		string(models.StatusConfigured): 0,
		string(models.StatusRunning):    0,
		string(models.StatusStopped):    0,
	}

	total := 0.0

	for _, ctx := range factory.Registry {
		state := string(ctx.Status)
		statusCounts[state]++
		total++
	}

	McpServerGauge.Set(total)

	for status, count := range statusCounts {
		McpServerState.WithLabelValues(status).Set(count)
	}
}
