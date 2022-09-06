package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Config struct {
	ServiceName     string `mapstructure:"METRIC_SERVICE_NAME"`
	ServiceHostPort string `mapstructure:"METRIC_SERVICE_HOST_PORT"`
}

type WalletServiceMetric struct {
	//CreateArticleHttpRequests  prometheus.Counter
	//CreateArticleKafkaRequests prometheus.Counter
	SuccessKafkaMessages prometheus.Counter
	ErrorKafkaMessages   prometheus.Counter
}

func NewWalletServiceMetric(cfg *Config) *WalletServiceMetric {
	return &WalletServiceMetric{
		//CreateArticleHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
		//	Name: fmt.Sprintf("%s_create_article_http_requests_total", cfg.ServiceName),
		//	Help: "The total of create article requests",
		//}),
		//CreateArticleKafkaRequests: promauto.NewCounter(prometheus.CounterOpts{
		//	Name: fmt.Sprintf("%s_create_article_kafka_requests_total", cfg.ServiceName),
		//	Help: "The total of create article kafka requests",
		//}),
		SuccessKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_kafka_processed_messages_total", cfg.ServiceName),
			Help: "The total number of success kafka processed messages",
		}),
		ErrorKafkaMessages: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_kafka_processed_messages_total", cfg.ServiceName),
			Help: "The total number of error kafka processed messages",
		}),
	}
}
