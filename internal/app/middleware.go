package app

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/mamun-jsx/Go-lang-Blog-Apps/config"
	"github.com/mamun-jsx/Go-lang-Blog-Apps/internal/auth"
	"gorm.io/gorm"
)

func InjectDB(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}

func JWTMiddleware(cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing Authorization"})
			}
			if !strings.HasPrefix(authHeader, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid authorization"})
			}
			token := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := auth.ValidateJWT(token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
			}
			c.Set("user_id", claims.ID)
			c.Set("role", claims.Role)
			return next(c)
		}
	}
}
