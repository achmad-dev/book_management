package util

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt"
)

// JWTCustomClaims defines custom claims for JWT
type JWTCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// NewAuthMiddleware creates a new JWT middleware
func NewAuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}

// GenerateToken generates a new JWT token
func GenerateToken(name, secret string) (string, error) {
	claims := &JWTCustomClaims{
		Username: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ValidateToken validates a JWT token
func ValidateToken(tokenString, secret string) (*JWTCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
