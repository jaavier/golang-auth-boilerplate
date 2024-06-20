package middlewares

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var jwtSecret []byte

func init() {
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
}

func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenStr := c.Request().Header.Get("Authorization")
		if tokenStr == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Missing token"})
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.ErrUnauthorized
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid token"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !claims["authorized"].(bool) {
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid token"})
		}

		c.Set("user_id", claims["user_id"])
		return next(c)
	}
}
