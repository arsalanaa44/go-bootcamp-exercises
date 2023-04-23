package userservice

import (
	"fmt"
	"gameapp/entity"
	"gameapp/pkg/phonenumber"
	//"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(user entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
	GetUserByID(userID int) (entity.User, error)
}

type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}

type Service struct {
	Repository Repository
	auth       AuthGenerator
}

func New(repository Repository, auth AuthGenerator) Service {
	return Service{repository, auth}
}

type RegisterRequest struct {
	Name        string `json:"name",xml:"Name"` // standard for each format
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type RegisterResponse struct {
	//User entity.User `json:"user"`
	User struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phoneNumber"`
	} `json:"user"`
}

func (s Service) Register(request RegisterRequest) (RegisterResponse, error) {
	// TODO - we should verify phone number by verification code

	// validate phone number
	if !phonenumber.IsValid(request.PhoneNumber) {

		return RegisterResponse{}, fmt.Errorf("phone number is not valid")
	}

	// check uniqueness of phone number
	if isUnique, iErr := s.Repository.IsPhoneNumberUnique(request.PhoneNumber); iErr != nil || !isUnique {
		if iErr != nil {

			return RegisterResponse{}, fmt.Errorf("unexpected error: %w", iErr)
		}
		if !isUnique {

			return RegisterResponse{}, fmt.Errorf("phone number is not unique")
		}
	}

	// validate name
	// seems there is no need to be function
	if len(request.Name) < 3 {

		return RegisterResponse{}, fmt.Errorf("name length should be greater than 3")
	}

	// TODO - check the password with regex pattern
	// validate password
	if len(request.Password) < 8 {

		return RegisterResponse{}, fmt.Errorf("password length should be greater than 8")
	}

	// create new user in storage
	var user entity.User
	if u, rErr := s.Repository.Register(entity.User{
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
		Password:    hashPassword(request.Password),
	}); rErr != nil {

		return RegisterResponse{}, fmt.Errorf("unexpected error: %w", rErr)
	} else {
		user = u
	}

	// return created user
	// TODO - assign anonymous struct
	var regRes RegisterResponse
	regRes.User.ID = user.ID
	regRes.User.Name = user.Name
	regRes.User.PhoneNumber = user.PhoneNumber

	return regRes, nil
}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	// TODO - it would be better to user have two separate method for getByPhoneNumber and checkExistence

	if user, exist, gErr := s.Repository.GetUserByPhoneNumber(req.PhoneNumber); gErr != nil {

		return LoginResponse{}, fmt.Errorf("enexpected error: %w", gErr)
	} else {
		if exist {
			if checkPasswordHash(req.Password, user.Password) {

				if aToken, cErr := s.auth.CreateAccessToken(user); cErr != nil {

					return LoginResponse{}, fmt.Errorf("enexpected error: %w", cErr)
				} else {

					if rToken, cErr := s.auth.CreateRefreshToken(user); cErr != nil {

						return LoginResponse{}, fmt.Errorf("enexpected error: %w", cErr)
					} else {

						return LoginResponse{aToken, rToken}, nil
					}
				}
			}
		}

		return LoginResponse{}, fmt.Errorf("phone-number or password is not correct")
	}

	// generate random session ID
	// save session ID in database
	// return session_ID to user
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type ProfileRequest struct {
	UserID int `json:"user_id"`
}

type ProfileResponse struct {
	Name string `json:"name"`
}

// all request services should be sanitized
func (s Service) Profile(req ProfileRequest) (ProfileResponse, error) {
	user, gErr := s.Repository.GetUserByID(req.UserID)
	if gErr != nil {
		// we assume input is sanitized/
		// TODO - we can use richError
		return ProfileResponse{}, fmt.Errorf("unexpected error: %w", gErr)
	}
	return ProfileResponse{Name: user.Name}, nil
}

//type Claims struct {
//	RegisteredClaims jwt.RegisteredClaims
//	UserID           int
//}
//
//func (c Claims) Valid() error {
//	return nil
//}
//func createToken(userID int, signKey string) (string, error) {
//	// create a signer for rsa 256
//	// TODO - rsa 256- https://github.com/golang-jwt/jwt/blob/main/http_example_test.go
//	//t := jwt.New(jwt.GetSigningMethod("RS256"))
//
//	// set our claims
//	claims := Claims{
//		jwt.RegisteredClaims{
//			// set the expire time
//			// see https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.4
//			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
//		},
//		userID,
//	}
//	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	if tokenString, err := accessToken.SigningString([]byte(signKey)); err != nil {
//
//		return "", err
//	} else {
//
//		return tokenString, nil
//	}
//}

//var jwtKey = []byte("your_secret_key")

//// Claims struct to define the JWT claims
//type Claims struct {
//	UserID             int `json:"user_id"`
//	jwt.StandardClaims     // it's not a field, inheritance, true json tags
//}
//
//// CreateToken function to create JWT token
//func createToken(userID int, jwtKey []byte) (string, error) {
//	// Define token expiration time
//	expirationTime := time.Now().Add(5 * time.Hour)
//
//	// Define token claims
//	claims := &Claims{
//		UserID: userID,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: expirationTime.Unix(),
//		},
//	}
//
//	// Create token using HS256 algorithm and the secret key
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	tokenString, err := token.SignedString(jwtKey)
//	if err != nil {
//		return "", err
//	}
//
//	return tokenString, nil
//}
