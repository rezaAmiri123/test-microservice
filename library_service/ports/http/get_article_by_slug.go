package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

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
