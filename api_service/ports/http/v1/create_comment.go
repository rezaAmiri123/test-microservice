package v1

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/api_service/dto"
	"github.com/rezaAmiri123/test-microservice/pkg/auth"
)

// CreateComment
// Create godoc
// @Summary Create comment
// @Description Create comment handler
// @Tags Comments
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Param article body dto.CreateCommentRequest true "New Comment"
// @Accept json
// @Produce json
// @Success 201 {object} dto.CreateCommentResponse
// @Router /comments/create [post]
func (h *HttpServer) CreateComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.metrics.CreateCommentHttpRequests.Inc()

		span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "HttpServer.CreateComment")
		defer span.Finish()

		u := auth.UserFromCtx(c.Request().Context())
		req := &dto.CreateCommentRequest{}
		if err := c.Bind(req); err != nil {
			h.log.WarnMsg("Bind", err)
			h.traceErr(span, err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		req.UUID = uuid.NewString()
		req.UserUUID = u.UUID
		if err := h.validate.StructCtx(ctx, req); err != nil {
			h.log.WarnMsg("validate", err)
			h.traceErr(span, err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := h.app.Commands.CreateComment.Handle(ctx, req); err != nil {
			h.log.WarnMsg("CreateComment", err)
			h.metrics.ErrorHttpRequests.Inc()
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		h.metrics.SuccessHttpRequests.Inc()
		return c.JSON(http.StatusOK, dto.CreateCommentResponse{UUID: req.UUID})
	}
}
