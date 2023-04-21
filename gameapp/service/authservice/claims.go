package authservice

import "github.com/dgrijalva/jwt-go"

// Claims struct to define the JWT claims
type Claims struct {
	UserID             int `json:"user_id"`
	jwt.StandardClaims     // it's not a field, inheritance, true json tags
}
