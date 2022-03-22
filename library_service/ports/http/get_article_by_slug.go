package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

func (h *HttpServer) GetBySlug() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.metric.GetArticleBySlugHttpRequests.Inc()

		span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "HttpServer.GetBySlug")
		defer span.Finish()

		articleSlug := c.Param("slug")
		a, err := h.app.Queries.GetArticleBySlug.Handle(ctx, articleSlug)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, a)
	}
}
