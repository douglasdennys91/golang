package jwt

import (
	"delivery-app/src/domain/entities"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type Claims struct {
	User entities.User `json:"user"`
	jwt.StandardClaims
}

func GenerateTokenJWT(data entities.User) (string, error) {
	expiration := time.Now().Add(1000 * time.Minute)
	claims := &Claims{
		User: data,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expiration.Unix()},
	}

	generate := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := generate.SignedString([]byte(os.Getenv("APP_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}

func IsAuthenticateJWT(data string) (*entities.User, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(data, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("APP_SECRET")), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}

	return &claims.User, nil
}