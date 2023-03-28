package task

import (
	"fmt"
	"todo_cli/entity"
)

type ServiceRepository interface {
	//DoesThisUserHaveThisCategoryID(userID, categoryID int) bool
	CreateNewTask(t entity.Task) (entity.Task, error)
	ListUserTasks(userID int) ([]entity.Task, error)
}
type Service struct {
	repository ServiceRepository
}

func NewService(repository ServiceRepository) Service {
	return Service{repository}
}

type CreateRequest struct {
	Title               string
	DueDate             string
	CategoryID          int
	AuthenticatedUserID int
}

type CreateResponse struct {
	Task entity.Task
}

func (t *Service) Create(req CreateRequest) (CreateResponse, error) {

	//if !t.repository.DoesThisUserHaveThisCategoryID(req.AuthenticatedUserID, req.CategoryID) {
	//	return CreateResponse{}, fmt.Errorf("category-ID %d is not valid", req.CategoryID)
	//}

	createdTask, cErr := t.repository.CreateNewTask(entity.Task{
		Title:      req.Title,
		DueDate:    req.DueDate,
		CategoryID: req.CategoryID,
		IsDone:     false,
		UserID:     req.AuthenticatedUserID,
	})
	if cErr != nil {
		return CreateResponse{}, fmt.Errorf("can't create new task", cErr)
	}

	return CreateResponse{createdTask}, nil

}

type ListRequest struct {
	AuthenticatedUserID int
}
type ListResponse struct {
	Tasks []entity.Task
}

func (t *Service) List(req ListRequest) (ListResponse, error) {
	if tasks, lErr := t.repository.ListUserTasks(req.AuthenticatedUserID); lErr != nil {
		return ListResponse{}, fmt.Errorf("can't list user tasks: %v", lErr)
	} else {
		return ListResponse{tasks}, nil
	}

}
