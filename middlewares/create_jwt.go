package middlewares

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func init() {
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
}

func CreateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
