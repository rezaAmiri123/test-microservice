package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/api_service/dto"
)

func (h *HttpServer) GetEmailByUUID() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.metrics.GetEmailByUUIDHttpRequests.Inc()

		span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "HttpServer.GetEmailByUUID")
		defer span.Finish()

		req := &dto.GetEmailByUUIDRequest{UUID: c.Param("uuid")}
		res, err := h.app.Queries.GetEmailByUUID.Handle(ctx, req)
		if err != nil {
			h.log.WarnMsg("GetEmailByUUID", err)
			h.metrics.ErrorHttpRequests.Inc()
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		h.metrics.SuccessHttpRequests.Inc()
		return c.JSON(http.StatusOK, res)
	}
}
