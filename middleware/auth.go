package middleware

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		storedPasswordHash := os.Getenv("HASHED_AUTH_PASSWORD")
		password := c.Request().Header.Get("X-Password")

		if isAuth := bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(password)); isAuth != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
		}
		return next(c)
	}
}
