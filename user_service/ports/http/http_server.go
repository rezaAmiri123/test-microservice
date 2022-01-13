package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rezaAmiri123/test-microservice/user_service/app"
	"github.com/rezaAmiri123/test-microservice/user_service/metrics"
	"net/http"
)

type HttpServer struct {
	app    *app.Application
	metric *metrics.UserServiceMetric
}

func NewHttpServer(addr string, application *app.Application, metric *metrics.UserServiceMetric) (*http.Server, error) {
	httpServer := &HttpServer{app: application, metric: metric}
	router := newRouter(httpServer)
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}, nil
}


func newRouter(httpServer *HttpServer) chi.Router {
	apiRouter := chi.NewRouter()
	setMiddlewares(apiRouter)
	apiRouter.Route("/users", func(r chi.Router) {
		r.Get("/profile", httpServer.GetProfile)
		r.Post("/register", httpServer.CreateUser)
		r.Post("/login", httpServer.Login)
	})

	rootRouter := chi.NewRouter()
	// we are mounting all APIs under /api path
	rootRouter.Mount("/api/v1", apiRouter)
	return rootRouter
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
