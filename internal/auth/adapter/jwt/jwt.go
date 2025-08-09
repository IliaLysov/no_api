package jwt

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWT struct {
	secret  string
	signer  jwt.SigningMethod
	keyFunc jwt.Keyfunc
}

func New() *JWT {
	secret := os.Getenv("JWT_SECRET")
	return &JWT{
		signer: jwt.SigningMethodHS256,
		secret: secret,
		keyFunc: func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		},
	}
}

func (j *JWT) CreateToken(_ context.Context, sub string) (string, error) {
	token := jwt.NewWithClaims(j.signer,
		jwt.MapClaims{
			"sub": sub,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		},
	)
	tokenString, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWT) Verify(tokenString string) (string, error) {
	parsed, err := jwt.Parse(tokenString, j.keyFunc)
	if err != nil || !parsed.Valid {
		return "", errors.New("invalid token")
	}
	claims := parsed.Claims.(jwt.MapClaims)
	sub, _ := claims["sub"].(string)
	return sub, nil
}
