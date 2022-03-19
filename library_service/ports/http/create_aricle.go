package http

import (
	"github.com/labstack/echo/v4"
	"github.com/rezaAmiri123/test-microservice/library_service/domain/article"
	"github.com/rezaAmiri123/test-microservice/pkg/auth"
	"net/http"
)

func (h *HttpServer) CreateArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		a := &article.Article{}
		//if err := json.NewDecoder(c.Request().Body).Decode(a); err != nil {
		//	// http.Error(w, err.Error(), http.StatusBadRequest)
		//	return c.JSON(http.StatusBadRequest, err.Error())
		//}
		if err := c.Bind(a); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
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
