package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

)

type Config struct {
	ServiceName string
	ServiceHostPort string
}

type UserServiceMetric struct {
	CreateUserHttpRequests prometheus.Counter
	SuccessKafkaMessages prometheus.Counter
	ErrorKafkaMessages   prometheus.Counter

}

func NewUserServiceMetric(cfg *Config) *UserServiceMetric  {
	return &UserServiceMetric{
		CreateUserHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_user_http_requests_total", cfg.ServiceName),
			Help: "The total of create user requests",
		}),
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