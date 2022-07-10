package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rezaAmiri123/test-microservice/library_service/app"
	"github.com/rezaAmiri123/test-microservice/library_service/metrics"
	"github.com/rezaAmiri123/test-microservice/pkg/auth"
)

type HttpServer struct {
	app        *app.Application
	metric     *metrics.LibraryServiceMetric
	authClient auth.AuthClient
}

func NewHttpServer(addr string, application *app.Application, metric *metrics.LibraryServiceMetric, authClient auth.AuthClient) (*echo.Echo, error) {
	httpServer := &HttpServer{app: application, metric: metric, authClient: authClient}
	//router := newEchoRouter(httpServer)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	v1 := e.Group("/api/v1")
	articleGroup := v1.Group("/articles")
	articleGroup.POST("/create", httpServer.CreateArticle())
	articleGroup.GET("/article/:slug", httpServer.GetBySlug())
	return e, nil

	//return &http.Server{
	//	Addr:    addr,
	//	Handler: router,
	//}, nil
}

//func newEchoRouter(httpServer *HttpServer) *echo.Echo {
//	e := echo.New()
//	e.Use(middleware.Logger())
//	e.Use(middleware.Recover())
//	v1 := e.Group("/api/v1")
//	articleGroup := v1.Group("/articles")
//	articleGroup.POST("/create", httpServer.CreateArticle())
//	articleGroup.GET("/article/:slug", httpServer.GetBySlug())
//	return e
//}

// func setMiddlewares(router *chi.Mux) {
// 	router.Use(middleware.RequestID)
// 	router.Use(middleware.RealIP)
// 	//router.Use(logs.NewStructuredLogger(logrus.StandardLogger()))
// 	router.Use(middleware.Recoverer)

// 	//addCorsMiddleware(router)
// 	//addAuthMiddleware(router)

// 	router.Use(
// 		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
// 		middleware.SetHeader("X-Frame-Options", "deny"),
// 	)
// 	router.Use(middleware.NoCache)
// }
