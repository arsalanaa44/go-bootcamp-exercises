package userservice

import (
	"fmt"
	"gameapp/entity"
	"gameapp/pkg/phonenumber"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(user entity.User) (entity.User, error)
}

type Service struct {
	Repository Repository
}

func New(repository Repository) Service {
	return Service{repository}
}

type RegisterRequest struct {
	Name        string
	PhoneNumber string
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

	// create new user in storage
	var user entity.User
	if u, rErr := s.Repository.Register(entity.User{
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
	}); rErr != nil {

		return RegisterResponse{}, fmt.Errorf("unexpected error: %w", rErr)
	} else {
		user = u
	}

	// return created user
	return RegisterResponse{user}, nil
}
