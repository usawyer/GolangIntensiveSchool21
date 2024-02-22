package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"time"
)

var jwtKey = []byte("SecretYouShouldHide")

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = "123456"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errors.Wrap(err, "error generated JWT token")
	}

	return tokenString, nil
}
