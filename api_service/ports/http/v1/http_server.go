package v1

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/api_service/app"
	"github.com/rezaAmiri123/test-microservice/api_service/metrics"
	"github.com/rezaAmiri123/test-microservice/pkg/auth"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
)

type HttpServer struct {
	app        *app.Application
	metrics    *metrics.ApiServiceMetric
	authClient auth.AuthClient
	validate   *validator.Validate
	log        logger.Logger
}

func NewHttpServer(
	application *app.Application,
	metrics *metrics.ApiServiceMetric,
	log logger.Logger,
) (*echo.Echo, error) {
	httpServer := &HttpServer{
		app:      application,
		metrics:  metrics,
		validate: validator.New(),
		log:      log,
	}
	//router := newEchoRouter(httpServer)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	v1 := e.Group("/api/v1")
	articleGroup := v1.Group("/users")
	articleGroup.POST("/create", httpServer.CreateUser())
	//articleGroup.GET("/article/:slug", httpServer.GetBySlug())
	return e, nil
}

func (h *HttpServer) traceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
	h.metrics.ErrorHttpRequests.Inc()
}
