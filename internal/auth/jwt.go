package auth

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID int, secret string, ttl time.Duration) (string, error){
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}