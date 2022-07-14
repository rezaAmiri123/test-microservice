package v1

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/api_service/app"
	"github.com/rezaAmiri123/test-microservice/api_service/metrics"
	"github.com/rezaAmiri123/test-microservice/pkg/auth"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
)

const (
	maxHeaderBytes = 1 << 20
	stackSize      = 1 << 10 // 1 KB
	bodyLimit      = "2M"
	readTimeout    = 15 * time.Second
	writeTimeout   = 15 * time.Second
	gzipLevel      = 5
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
	authClient auth.AuthClient,
) (*echo.Echo, error) {
	httpServer := &HttpServer{
		app:        application,
		metrics:    metrics,
		validate:   validator.New(),
		log:        log,
		authClient: authClient,
	}
	//router := newEchoRouter(httpServer)
	e := echo.New()

	e.Server.ReadTimeout = readTimeout
	e.Server.WriteTimeout = writeTimeout
	e.Server.MaxHeaderBytes = maxHeaderBytes

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         stackSize,
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	e.Use(middleware.RequestID())
	e.Use(middleware.BodyLimit(bodyLimit))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestID},
	}))

	v1 := e.Group("/api/v1")

	userGroup := v1.Group("/users")
	userGroup.POST("/register", httpServer.CreateUser())
	userGroup.POST("/login", httpServer.UserLogin())

	articleGroup := v1.Group("/articles")
	articleGroup.POST("/create", httpServer.CreateArticle())
	articleGroup.GET("/list", httpServer.GetArticles())
	articleGroup.GET("/article/:slug", httpServer.GetArticleBySlug())
	return e, nil
}

func (h *HttpServer) traceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
	h.metrics.ErrorHttpRequests.Inc()
}
