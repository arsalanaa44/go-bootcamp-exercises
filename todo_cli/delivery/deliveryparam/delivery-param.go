package deliveryparam

import (
	"todo_cli/service/task"
	"todo_cli/service/user"
)

type ClientRequest struct {
	Command           string
	CreateTaskRequest task.CreateRequest
	LoginRequest      user.LoginRequest
}

type ClientResponse struct {
	Data string
}
