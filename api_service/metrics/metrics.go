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

type ApiServiceMetric struct {
	CreateUserHttpRequests       prometheus.Counter
	UserLoginHttpRequests        prometheus.Counter
	CreateArticleHttpRequests    prometheus.Counter
	GetArticleBySlugHttpRequests prometheus.Counter
	GetArticlesHttpRequests      prometheus.Counter
	SuccessHttpRequests          prometheus.Counter
	ErrorHttpRequests            prometheus.Counter
}

func NewApiServiceMetric(cfg *Config) *ApiServiceMetric {
	return &ApiServiceMetric{
		CreateUserHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_user_http_requests_total", cfg.ServiceName),
			Help: "The total of create user requests",
		}),
		UserLoginHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_user_login_http_requests_total", cfg.ServiceName),
			Help: "The total of user login requests",
		}),
		CreateArticleHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_create_article_http_requests_total", cfg.ServiceName),
			Help: "The total of create article requests",
		}),
		GetArticleBySlugHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_get_article_by_slug_http_requests_total", cfg.ServiceName),
			Help: "The total of get article by slug http requests",
		}),
		GetArticlesHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_get_articles_http_requests_total", cfg.ServiceName),
			Help: "The total of get articles http requests",
		}),
		SuccessHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_success_http_requsts_total", cfg.ServiceName),
			Help: "The total number of success http requests",
		}),
		ErrorHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_error_http_requsts_total", cfg.ServiceName),
			Help: "The total number of error http requests",
		}),
	}
}
