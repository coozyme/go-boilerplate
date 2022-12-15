package middleware

import (
	"fmt"
	"go-mitra-akosta/pkg/utils/config"
	"go-mitra-akosta/pkg/utils/logger"
	"go-mitra-akosta/pkg/utils/response"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type customContext struct {
	echo.Context
}

func AuthorizeJWT(cfg *config.Config, log *logger.Logging, next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return response.ResponseError(http.StatusUnauthorized, "unauthorized", 401).Send(c)
		}
		if !strings.Contains(authHeader, "Bearer") {
			return response.ResponseError(http.StatusUnauthorized, "unauthorized", 401).Send(c)
		}

		splitToken := strings.Split(authHeader, "Bearer ")

		token, error := jwt.Parse(splitToken[1], func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(cfg.Jwt.Secret), nil
		})

		if !token.Valid || error != nil {
			log.InfoLog("Token Invalid:", error)
			return response.ResponseError(http.StatusUnauthorized, "unauthorized", 401).Send(c)
		}

		claims := token.Claims.(jwt.MapClaims)

		c.Set("mitra", claims["mitra"])

		return next(c)
	}
}
