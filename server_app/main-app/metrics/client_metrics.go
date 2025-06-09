package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	McpClientGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mcp",
			Subsystem: "client",
			Name:      "client_ops_total",
			Help:      "Total number of active MCP clients",
		},
	)

	McpClientState = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "mcp",
			Subsystem: "client",
			Name:      "client_state_count",
			Help:      "MCP client counts by state",
		},
		[]string{"state"},
	)

	clientGaugeVal float64
)

func IncMcpClientGauge() {
	clientGaugeVal++
	McpClientGauge.Set(clientGaugeVal)
}

func DecMcpClientGauge() {
	if clientGaugeVal > 0 {
		clientGaugeVal--
	}
	McpClientGauge.Set(clientGaugeVal)
}

func McpClientGaugeValue() float64 {
	return clientGaugeVal
}

func SetMcpClientState(state string, val float64) {
	McpClientState.WithLabelValues(state).Set(val)
}

