package authentication

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/j-monteiro/movies-api/domain/models"
)

func Authenticate(token string) models.User {
	claims := jwt.MapClaims{}

	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	return models.User{}
}
