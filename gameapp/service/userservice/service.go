package userservice

import (
	"fmt"
	"gameapp/entity"
	"gameapp/pkg/phonenumber"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(user entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
}

type Service struct {
	Repository Repository
}

func New(repository Repository) Service {
	return Service{repository}
}

type RegisterRequest struct {
	Name        string `json:"name",xml:"Name"` // standard for each format
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type RegisterResponse struct {
	User entity.User
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
	return RegisterResponse{user}, nil
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
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	// TODO - it would be better to user have two separate method for getByPhoneNumber and checkExistence

	if user, exist, gErr := s.Repository.GetUserByPhoneNumber(req.PhoneNumber); gErr != nil {

		return LoginResponse{}, fmt.Errorf("enexpected error: %w", gErr)
	} else {
		if exist {
			if checkPasswordHash(req.Password, user.Password) {

				return LoginResponse{}, nil
			}
		}

		return LoginResponse{}, fmt.Errorf("phone-number or password is not correct")
	}

	return LoginResponse{}, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
