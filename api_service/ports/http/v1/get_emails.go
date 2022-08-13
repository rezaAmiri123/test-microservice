package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/pkg/pagnation"
)

func (h *HttpServer) GetEmails() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.metrics.GetEmailsHttpRequests.Inc()

		span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "HttpServer.GetEmails")
		defer span.Finish()

		query := pagnation.NewPaginationFromQueryParams(c.QueryParam("size"), c.QueryParam("page"))

		res, err := h.app.Queries.GetEmails.Handle(ctx, query)
		if err != nil {
			h.log.WarnMsg("GetEmails", err)
			h.metrics.ErrorHttpRequests.Inc()
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		h.metrics.SuccessHttpRequests.Inc()
		return c.JSON(http.StatusOK, res)
	}
}
