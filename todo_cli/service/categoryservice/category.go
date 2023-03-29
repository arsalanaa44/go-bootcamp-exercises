package categoryservice

import (
	"fmt"
	"todo_cli/entity"
)

type ServiceCategoryRepository interface {
	CreateNewCategory(category entity.Category) (entity.Category, error)
	UserIDAndCategoryIDValidation(userID, categoryID int) bool
}

type Service struct {
	repository ServiceCategoryRepository
}

func NewService(repository ServiceCategoryRepository) Service {

	return Service{repository}
}

type CreateRequest struct {
	Title               string
	Color               string
	AuthenticatedUserID int
}

type CreateResponse struct {
	Category entity.Category
}

func (s *Service) Create(request CreateRequest) (CreateResponse, error) {
	category, cErr := s.repository.CreateNewCategory(entity.Category{
		Title:  request.Title,
		Color:  request.Color,
		UserID: request.AuthenticatedUserID,
	})
	if cErr != nil {

		return CreateResponse{}, fmt.Errorf("can't create new categoryservice :%v", cErr)
	}

	return CreateResponse{category}, nil
}

func (s Service) DoesThisUserHaveThisCategoryID(userID, categoryID int) bool {

	return s.repository.UserIDAndCategoryIDValidation(userID, categoryID)
}
