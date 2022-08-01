package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Config struct {
	ServiceName     string
	ServiceHostPort string
}

type LibraryServiceMetric struct {
	CreateArticleHttpRequests    prometheus.Counter
	GetArticleBySlugHttpRequests prometheus.Counter
	GetArticleBySlugGrpcRequests prometheus.Counter
	GetArticlesGrpcRequests      prometheus.Counter
	CreateArticleKafkaRequests   prometheus.Counter
	CreateCommentKafkaRequests   prometheus.Counter
	SuccessKafkaMessages         prometheus.Counter
	ErrorKafkaMessages           prometheus.Counter
}

func NewLibraryServiceMetric(cfg *Config) *LibraryServiceMetric {
	return &LibraryServiceMetric{
		CreateArticleHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_article_http_requests_total", cfg.ServiceName),
			Help: "The total of create article requests",
		}),
		GetArticleBySlugHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_get_article_by_slug_http_requests_total", cfg.ServiceName),
			Help: "The total of get article by slug http requests",
		}),
		GetArticleBySlugGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_get_article_by_slug_grpc_requests_total", cfg.ServiceName),
			Help: "The total of get article by slug grpc requests",
		}),
		GetArticlesGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_get_articles_grpc_requests_total", cfg.ServiceName),
			Help: "The total of get articles grpc requests",
		}),
		CreateArticleKafkaRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_article_kafka_requests_total", cfg.ServiceName),
			Help: "The total of create article kafka requests",
		}),
		CreateCommentKafkaRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_comment_kafka_requests_total", cfg.ServiceName),
			Help: "The total of create comment kafka requests",
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
