package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	McpRegistryGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "mcp",
			Subsystem: "registry",
			Name:      "registry_ops_total",
			Help:      "Total number of defined MCP registries",
		},
	)

	McpRegistryState = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "mcp",
			Subsystem: "registry",
			Name:      "registry_state_count",
			Help:      "MCP registry counts by state",
		},
		[]string{"state"},
	)

	regGaugeVal float64
)

func IncMcpRegistryGauge() {
	regGaugeVal++
	McpRegistryGauge.Set(regGaugeVal)
}

func DecMcpRegistryGauge() {
	if regGaugeVal > 0 {
		regGaugeVal--
	}
	McpRegistryGauge.Set(regGaugeVal)
}

func McpRegistryGaugeValue() float64 {
	return regGaugeVal
}

func SetMcpRegistryState(state string, val float64) {
	McpRegistryState.WithLabelValues(state).Set(val)
}

