package v1

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/api_service/dto"
	"net/http"
)

func (h *HttpServer) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.metrics.CreateUserHttpRequests.Inc()

		span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "HttpServer.CreateUser")
		defer span.Finish()

		req := &dto.CreateUserRequest{}
		if err := c.Bind(req); err != nil {
			h.log.WarnMsg("Bind", err)
			h.traceErr(span, err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		req.UserID = uuid.NewString()
		if err := h.validate.StructCtx(ctx, req); err != nil {
			h.log.WarnMsg("validate", err)
			h.traceErr(span, err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := h.app.Commands.CreateUser.Handle(ctx, req); err != nil {
			h.log.WarnMsg("CreateUser", err)
			h.metrics.ErrorHttpRequests.Inc()
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		h.metrics.SuccessHttpRequests.Inc()
		return c.JSON(http.StatusOK, dto.CreateUserResponse{UserID: req.UserID})
	}
}
