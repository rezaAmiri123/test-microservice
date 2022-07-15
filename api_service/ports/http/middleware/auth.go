package middleware

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rezaAmiri123/test-microservice/pkg/auth"
	"github.com/rezaAmiri123/test-microservice/pkg/utils"
)

func (mw *MiddlewareManager) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := mw.authClient.VerifyToken(c.Request().Context(), auth.GetTokenFromHeader(c.Request()))
		if err != nil {
			mw.logger.Errorf("GetSessionByID RequestID: %s, Error: %s",
				utils.GetRequestID(c),
				err.Error(),
			)
			return c.JSON(http.StatusUnauthorized, auth.NoUserInContextError)
		}

		c.Set("user", user)

		ctx := context.WithValue(c.Request().Context(), auth.UserContextKey, user)
		c.SetRequest(c.Request().WithContext(ctx))

		mw.logger.Info(
			"AuthMiddleware, RequestID: %s,  IP: %s, UserID: %s",
			utils.GetRequestID(c),
			utils.GetIPAddress(c),
			user.UUID,
		)

		return next(c)
	}
}
