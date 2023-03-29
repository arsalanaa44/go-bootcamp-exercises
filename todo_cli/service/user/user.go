package user

import (
	"fmt"
	"todo_cli/contract"
	"todo_cli/encryption"
	"todo_cli/entity"
)

type ServiceTempRepository interface {
	Save(user entity.User) entity.User
	DeleteLast()
	ListUsers() []entity.User
}

type ServicePermRepository interface {
	contract.UserWriteStore
}

type Service struct {
	tempRepo ServiceTempRepository
	permRepo ServicePermRepository
}

func NewUserService(tempRepo ServiceTempRepository, permRepo ServicePermRepository) Service {
	return Service{tempRepo, permRepo}
}

type RegisterRequest struct {
	Name     string
	Email    string
	Password string
}
type RegisterResponse struct {
	User entity.User
}

func (s *Service) Register(req RegisterRequest) (RegisterResponse, error) {

	user := s.tempRepo.Save(entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if sErr := s.permRepo.Save(user); sErr != nil {
		s.tempRepo.DeleteLast()

		return RegisterResponse{}, fmt.Errorf("can't save it permanantly ,%v", sErr)
	}

	return RegisterResponse{user}, nil

}

type ListRequest struct {
}
type ListResponse struct {
	Users []entity.User
}

func (s *Service) List() ListResponse {
	return ListResponse{
		s.tempRepo.ListUsers()}
}

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	User *entity.User
}

func (s *Service) Login(request LoginRequest) (LoginResponse, error) {
	for _, user := range s.List().Users {
		if user.Email == request.Email {
			er := encryption.PassValidation([]byte(user.Password), []byte(request.Password))
			if er == nil {

				return LoginResponse{&user}, nil
			}
		}
	}

	return LoginResponse{}, fmt.Errorf("email or password is incorrect")
}
