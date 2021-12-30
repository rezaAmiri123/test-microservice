package domain

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"
)

var jwtSecret = []byte("123456789")

type claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) GenerateJWTToken() (string, error) {
	claims := &claims{
		u.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3 * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return t, nil
}
func GetUsernameFromJWTToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if !token.Valid {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return "", errors.New("invalid token: it's not even a token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return "", errors.New("token expired")
			} else {
				return "", fmt.Errorf("invalid token: couldn't handle this token; %w", err)
			}
		} else {
			return "", fmt.Errorf("invalid token: couldn't handle this token; %w", err)
		}
	}
	c, ok := token.Claims.(*claims)
	if !ok {
		return "", errors.New("invalid token: cannot map token to claims")
	}
	if c.ExpiresAt < time.Now().Unix() {
		return "", errors.New("token expired")
	}
	return c.Username, nil
}
