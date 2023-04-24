package authservice

import (
	"fmt"
	"gameapp/entity"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type Config struct {
	SignKey               string
	AccessExpirationTime  time.Duration
	RefreshExpirationTime time.Duration
	AccessSubject         string
	RefreshSubject        string
}

type Service struct {
	config Config
}

func New(cfg Config) Service {
	return Service{cfg}
}

func (s Service) CreateAccessToken(user entity.User) (string, error) {

	return createToken(user.ID, []byte(s.config.SignKey), s.config.AccessSubject, s.config.AccessExpirationTime)
}

func (s Service) CreateRefreshToken(user entity.User) (string, error) {

	return createToken(user.ID, []byte(s.config.SignKey), s.config.RefreshSubject, s.config.RefreshExpirationTime)
}

func (s Service) ParseToken(tokenString string) (*Claims, error) {

	tokenString = strings.Replace(tokenString, "Bearer ", "", -1)
	// Parse token to extract claims
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(s.config.SignKey), nil
	})
	if err != nil {
		return nil, err
	}

	// Extract claims from token
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

// CreateToken function to create JWT token
func createToken(userID int, jwtKey []byte, subject string, expireAt time.Duration) (string, error) {
	// Define token expiration time
	expirationTime := time.Now().Add(expireAt)

	// Define token claims
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			Subject:   subject,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create token using HS256 algorithm and the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//func verifyAuthHeader(authHeader string) (string, error) {
//	// Get the Authorization header value
//
//	if authHeader == "" {
//
//		return "", fmt.Errorf("401, you did not set a proper token")
//	}
//
//	// Check that the Authorization header is in the Bearer format
//	authHeaderParts := strings.Split(authHeader, " ")
//	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
//		return "", fmt.Errorf("401, you did not set a proper token")
//	}
//
//	// Extract the JWT token string
//	return authHeaderParts[1], nil
//
//	// The token is valid and the user is authorized, so continue with the handler logic...
//}
