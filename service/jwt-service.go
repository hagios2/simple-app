package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type JWTService interface {
	GenerateToken(name string, admin bool) string
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "hagioswilson",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtService) GenerateToken(username string, admin bool) string {
	// Set custom and standard claims
	claims := &jwtCustomClaims{
		Name:  username,
		Admin: admin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * 72)},
			Issuer:    service.issuer,
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token using the secret signing key
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t

}

func (service *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			// Signing method validation
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// Return the secret signing key
			return []byte(service.secretKey), nil
		},
	)
}
