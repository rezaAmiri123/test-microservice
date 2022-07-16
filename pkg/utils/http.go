package utils

import (
	"context"

	"github.com/labstack/echo/v4"
)

// ReqIDCtxKey is a key used for the Request ID in context
type ReqIDCtxKey struct{}

// Get request id from echo context
func GetRequestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}

// Get user ip address
func GetIPAddress(c echo.Context) string {
	return c.Request().RemoteAddr
}

// Get context  with request id
func GetRequestCtx(c echo.Context) context.Context {
	return context.WithValue(c.Request().Context(), ReqIDCtxKey{}, GetRequestID(c))
}
