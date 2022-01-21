package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/labstack/echo/v4"
	"github.com/rezaAmiri123/test-microservice/library_service/app"
	"github.com/rezaAmiri123/test-microservice/library_service/metrics"
	"github.com/rezaAmiri123/test-microservice/pkg/auth"
)

type HttpServer struct {
	app        *app.Application
	metric     *metrics.ArticleServiceMetric
	authClient auth.AuthClient
}

func NewHttpServer(addr string, application *app.Application, metric *metrics.ArticleServiceMetric, authClient auth.AuthClient) (*http.Server, error) {
	httpServer := &HttpServer{app: application, metric: metric, authClient: authClient}
	router := newEchoRouter(httpServer)
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}, nil
}

// func newRouter(httpServer *HttpServer) chi.Router {
// 	apiRouter := chi.NewRouter()
// 	setMiddlewares(apiRouter)
// 	apiRouter.Route("/users", func(r chi.Router) {
// 		r.Get("/profile", httpServer.GetProfile)
// 		r.Post("/register", httpServer.CreateUser)
// 		r.Post("/login", httpServer.Login)
// 	})

// 	rootRouter := chi.NewRouter()
// 	// we are mounting all APIs under /api path
// 	rootRouter.Mount("/api/v1", apiRouter)
// 	return rootRouter
// }

func newEchoRouter(httpServer *HttpServer) *echo.Echo {
	e := echo.New()
	v1 := e.Group("/api/v1")
	articleGroup := v1.Group("/articles")
	articleGroup.POST("/create", httpServer.CreateArticle())
	return e
}

func setMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	//router.Use(logs.NewStructuredLogger(logrus.StandardLogger()))
	router.Use(middleware.Recoverer)

	//addCorsMiddleware(router)
	//addAuthMiddleware(router)

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)
}
