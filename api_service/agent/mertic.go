package agent

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rezaAmiri123/test-microservice/api_service/metrics"
	"net/http"
)

func (a *Agent) setupMetric() error {
	metric := metrics.NewApiServiceMetric(&a.MetricConfig)
	//prometheus.MustRegister(metric.CreateUserHttpRequests)
	a.metric = metric
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(a.MetricConfig.ServiceHostPort, nil)
	return nil
}
