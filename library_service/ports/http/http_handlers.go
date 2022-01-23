package http

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaAmiri123/test-microservice/library_service/domain/article"
	"github.com/rezaAmiri123/test-microservice/pkg/auth"
)

func (h *HttpServer) CreateArticle1(w http.ResponseWriter, r *http.Request) {
	a := &article.Article{}
	if err := json.NewDecoder(r.Body).Decode(a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u, err := auth.UserFromCtx(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.app.Commands.CreateArticle.Handle(r.Context(), a, u.UUID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *HttpServer) CreateArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		a := &article.Article{}
		if err := json.NewDecoder(c.Request().Body).Decode(a); err != nil {
			// http.Error(w, err.Error(), http.StatusBadRequest)
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		u, err := h.authClient.VerityToken(c.Request().Context(), auth.GetTokenFromHeader(c.Request()))
		if err != nil {
			// http.Error(w, err.Error(), http.StatusBadRequest)
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if err := h.app.Commands.CreateArticle.Handle(c.Request().Context(), a, u.UUID); err != nil {
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		// w.WriteHeader(http.StatusOK)
		h.metric.CreateArticleHttpRequests.Inc()
		return c.JSON(http.StatusOK, nil)
	}
}

//func (h *HttpServer) GetArticle(w http.ResponseWriter, r *http.Request) {
//	articleSlug := chi.URLParam(r, "slug")
//	a, err := h.app.Queries.GetArticle.Handle(r.Context(), articleSlug)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	err = json.NewEncoder(w).Encode(a)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//}

func (h *HttpServer) GetBySlug() echo.HandlerFunc {
	return func(c echo.Context) error {
		articleSlug := c.Param("slug")
		a, err := h.app.Queries.GetArticleBySlug.Handle(c.Request().Context(), articleSlug)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, a)
	}
}

//func (h newsHandlers) GetByID() echo.HandlerFunc {
//	return func(c echo.Context) error {
//		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "newsHandlers.GetByID")
//		defer span.Finish()
//
//		newsUUID, err := uuid.Parse(c.Param("news_id"))
//		if err != nil {
//			utils.LogResponseError(c, h.logger, err)
//			return c.JSON(httpErrors.ErrorResponse(err))
//		}
//
//		newsByID, err := h.newsUC.GetNewsByID(ctx, newsUUID)
//		if err != nil {
//			utils.LogResponseError(c, h.logger, err)
//			return c.JSON(httpErrors.ErrorResponse(err))
//		}
//
//		return c.JSON(http.StatusOK, newsByID)
//	}
//}
