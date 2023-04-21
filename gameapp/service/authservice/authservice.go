package authservice

import (
	"fmt"
	"gameapp/entity"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Service struct {
	signKey               string
	accessExpirationTime  time.Duration
	refreshExpirationTime time.Duration
	accessSubject         string
	refreshSubject        string
}

func New(signKey, accessSubject, refreshSubject string,
	accessExpirationTime, refreshExpirationTime time.Duration) Service {
	return Service{
		signKey,
		accessExpirationTime,
		refreshExpirationTime,
		accessSubject,
		refreshSubject,
	}
}
func (s Service) CreateAccessToken(user entity.User) (string, error) {

	return createToken(user.ID, []byte(s.signKey), s.accessSubject, s.accessExpirationTime)
}

func (s Service) CreateRefreshToken(user entity.User) (string, error) {

	return createToken(user.ID, []byte(s.signKey), s.refreshSubject, s.refreshExpirationTime)
}

func (s Service) ParseToken(tokenString string) (*Claims, error) {

	// Parse token to extract claims
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(s.signKey), nil
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
