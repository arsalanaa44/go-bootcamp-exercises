package taskclient

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"todo_cli/service/task"
)

type ClientService struct {
}

func NewClientService() *ClientService {
	return &ClientService{}
}

func (cs *ClientService) CreateTask() (task.CreateRequest, error) {

	scanner := bufio.NewScanner(os.Stdin)
	var title, dueDate, category string

	fmt.Println("please enter the taskclient title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the taskclient due date")
	scanner.Scan()
	dueDate = scanner.Text()

	fmt.Println("please enter the taskclient categoryservice-ID")
	scanner.Scan()
	category = scanner.Text()
	categoryID, e := strconv.Atoi(category)
	if e != nil {

		return task.CreateRequest{}, fmt.Errorf("categoryservice-ID is not valid integer : %v", e)
	}
	return task.CreateRequest{
		Title:      title,
		DueDate:    dueDate,
		CategoryID: categoryID,
	}, nil
}
