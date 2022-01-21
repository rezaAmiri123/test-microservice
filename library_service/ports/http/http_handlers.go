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
		return c.JSON(http.StatusOK, nil)
	}
}
