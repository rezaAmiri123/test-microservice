package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/api_service/dto"
)

func (h *HttpServer) GetArticleBySlug() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.metrics.GetArticleBySlugHttpRequests.Inc()

		span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "HttpServer.GetArticleBySlug")
		defer span.Finish()

		req := &dto.GetArticleBySlugRequest{Slug: c.Param("slug")}
		res, err := h.app.Queries.GetArticleBySlug.Handle(ctx, req)
		if err != nil {
			h.log.WarnMsg("GetArticleBySlug", err)
			h.metrics.ErrorHttpRequests.Inc()
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		h.metrics.SuccessHttpRequests.Inc()
		return c.JSON(http.StatusOK, res)
	}
}
