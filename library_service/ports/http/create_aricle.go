package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/library_service/ports/dto"
	"github.com/rezaAmiri123/test-microservice/pkg/auth"
)

func (h *HttpServer) CreateArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.metric.CreateArticleHttpRequests.Inc()

		span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "HttpServer.CreateArticle")
		defer span.Finish()

		req := &dto.CreateArticleRequest{}
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := req.Validate(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		u, err := h.authClient.VerityToken(ctx, auth.GetTokenFromHeader(c.Request()))
		if err != nil {
			// http.Error(w, err.Error(), http.StatusBadRequest)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := h.app.Commands.CreateArticle.Handle(ctx, req.MapToArticle(), u.UUID); err != nil {
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, nil)
	}
}
