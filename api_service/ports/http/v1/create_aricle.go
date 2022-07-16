package v1

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/api_service/dto"
	"github.com/rezaAmiri123/test-microservice/pkg/auth"
)

// CreateArticle
// Create godoc
// @Summary Create article
// @Description Create article handler
// @Tags Articles
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Param article body dto.CreateArticleRequest true "New Article"
// @Accept json
// @Produce json
// @Success 201 {object} dto.CreateArticleResponse
// @Router /articles/create [post]
func (h *HttpServer) CreateArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.metrics.CreateArticleHttpRequests.Inc()

		span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "HttpServer.CreateArticle")
		defer span.Finish()

		u := auth.UserFromCtx(c.Request().Context())
		req := &dto.CreateArticleRequest{}
		if err := c.Bind(req); err != nil {
			h.log.WarnMsg("Bind", err)
			h.traceErr(span, err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		req.ArticleID = uuid.NewString()
		req.UserID = u.UUID
		if err := h.validate.StructCtx(ctx, req); err != nil {
			h.log.WarnMsg("validate", err)
			h.traceErr(span, err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := h.app.Commands.CreateArticle.Handle(ctx, req); err != nil {
			h.log.WarnMsg("CreateArticle", err)
			h.metrics.ErrorHttpRequests.Inc()
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		h.metrics.SuccessHttpRequests.Inc()
		return c.JSON(http.StatusOK, dto.CreateArticleResponse{ArticleID: req.ArticleID})
	}
}
