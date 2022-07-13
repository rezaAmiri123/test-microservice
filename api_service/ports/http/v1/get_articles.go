package v1

import (
	"github.com/rezaAmiri123/test-microservice/pkg/pagnation"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

func (h *HttpServer) GetArticles() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.metrics.GetArticlesHttpRequests.Inc()

		span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "HttpServer.GetArticles")
		defer span.Finish()

		query := pagnation.NewPaginationFromQueryParams(c.QueryParam("size"), c.QueryParam("page"))

		res, err := h.app.Queries.GetArticles.Handle(ctx, query)
		if err != nil {
			h.log.WarnMsg("GetArticles", err)
			h.metrics.ErrorHttpRequests.Inc()
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		h.metrics.SuccessHttpRequests.Inc()
		return c.JSON(http.StatusOK, res)
	}
}
