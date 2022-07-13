package v1

import (
	"github.com/google/uuid"
	"github.com/rezaAmiri123/test-microservice/api_service/dto"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/test-microservice/pkg/auth"
)

func (h *HttpServer) CreateArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.metrics.CreateArticleHttpRequests.Inc()

		span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "HttpServer.CreateArticle")
		defer span.Finish()

		u, err := h.authClient.VerifyToken(ctx, auth.GetTokenFromHeader(c.Request()))
		if err != nil {
			h.log.WarnMsg("verify token", err)
			h.traceErr(span, err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

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
