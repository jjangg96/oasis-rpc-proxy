package metric

import (
	"fmt"
	"github.com/jjangg96/oasis-rpc-proxy/utils/logger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	ClientRequestDuration = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "figment",
			Subsystem: "oasis_proxy",
			Name:      "node_request_duration",
			Help:      "The total time required to execute request to node",
		},
		[]string{"request"},
	)
)

// ClientMetric handles HTTP requests
type ClientMetric struct{}

// NewIndexerMetric returns a new server instance
func NewClientMetric() *ClientMetric {
	app := &ClientMetric{}
	return app.init()
}

func (m *ClientMetric) StartServer(listenAdd string, url string) error {
	logger.Info(fmt.Sprintf("starting metric server at %s...", url), logger.Field("app", "server"))

	http.Handle(url, promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: true,
		},
	))
	return http.ListenAndServe(listenAdd, nil)
}

func (m *ClientMetric) init() *ClientMetric {
	logger.Info("initializing metric server...", logger.Field("app", "server"))

	prometheus.MustRegister(ClientRequestDuration)

	// Add Go module build info.
	prometheus.MustRegister(prometheus.NewBuildInfoCollector())
	return m
}
