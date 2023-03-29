package userclient

import (
	"bufio"
	"fmt"
	"os"
	"todo_cli/service/user"
)

type ClientService struct {
}

func NewClientService() *ClientService {
	return &ClientService{}
}

func (cs *ClientService) Login() user.LoginRequest {
	fmt.Println("login process :")
	scanner := bufio.NewScanner(os.Stdin)
	var email, password string

	fmt.Println("please enter the user email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the user password")
	scanner.Scan()
	password = scanner.Text()
	return user.LoginRequest{email, password}
}
