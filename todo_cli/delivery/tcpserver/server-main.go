package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"todo_cli/constant"
	"todo_cli/delivery/deliveryparam"
	"todo_cli/entity"
	"todo_cli/repository/filestore"
	"todo_cli/repository/inmemory"
	"todo_cli/service/categoryservice"
	"todo_cli/service/task"
	"todo_cli/service/user"
)

var (
	authenticatedUser *entity.User
	serializationMode string
	cRequest          = deliveryparam.ClientRequest{}
	cResponse         = deliveryparam.ClientResponse{}
)

const (
	userStoragePath = "user.txt"
)

func main() {

	taskMemoryRepo := inmemory.NewTaskStore()
	categoryMemoryRepo := inmemory.NewCategoryStore()
	categoryService := categoryservice.NewService(categoryMemoryRepo)
	taskService := task.NewService(taskMemoryRepo, categoryService)

	fmt.Println("hello to TODO app")

	serializeMode := flag.String("serialization-mode", constant.JsonSerializationMode, "serialization mode to write data to file")
	flag.Parse()

	switch *serializeMode {
	case constant.MandaravardiSerializationMode:
		{
			serializationMode = constant.MandaravardiSerializationMode
		}
	default:
		{
			serializationMode = constant.JsonSerializationMode
		}
	}
	var userFileStore = filestore.New(userStoragePath, serializationMode)
	userMemoryStore := inmemory.NewUserStore(userFileStore.Load())
	userService := user.NewUserService(userMemoryStore, userFileStore)

	const (
		network = "tcp"
		address = "127.0.0.1:2022"
	)

	// create listener
	var listener net.Listener
	if l, lErr := net.Listen(network, address); lErr != nil {
		fmt.Println("lError :", lErr)
	} else {
		listener = l
		defer listener.Close()
	}

	fmt.Println("listener address :", listener.Addr())

	// listen for new connection
	var connection net.Conn
	for {
		cRequest = deliveryparam.ClientRequest{}
		cResponse = deliveryparam.ClientResponse{}

		if c, aErr := listener.Accept(); aErr != nil {
			log.Println("aError :", aErr)

			continue
		} else {
			connection = c
		}

		fmt.Println("client address:", connection.RemoteAddr())

		// process request
		var data = make([]byte, 1024)
		if numberOfReadBytes, rErr := connection.Read(data); rErr != nil {

			cResponse.Data = fmt.Sprintln("reading error", rErr)

			continue
		} else {
			if jErr := json.Unmarshal(data[:numberOfReadBytes], &cRequest); jErr != nil {

				fmt.Println(string(data))
				cResponse.Data = fmt.Sprintln("unmarshal error", jErr)
			} else {

				runCommand(&userService, &taskService)
			}
		}
		message, _ := json.Marshal(cResponse)
		fmt.Println(string(message))
		if _, wErr := connection.Write(message); wErr != nil {
			log.Println("rError", wErr)

			continue
		}
		connection.Close()
		fmt.Println(authenticatedUser)
		if authenticatedUser != nil {
			fmt.Println(taskService.List(task.ListRequest{authenticatedUser.ID}))
		}
	}
}

func runCommand(userService *user.Service, taskService *task.Service) {

	switch cRequest.Command {
	case "login":
		{

			login(userService)
		}
	case "create-task":
		{

			createTask(taskService)
		}
	case "exit":
		{
			os.Exit(0)
		}
	}
}

func login(userService *user.Service) {

	loginRequest := cRequest.LoginRequest
	//autologin
	//loginRequest := user.LoginRequest{
	//	Email:    "pp",
	//	Password: "ppp",
	//}

	if loginResponse, lErr := userService.Login(loginRequest); lErr != nil {

		cResponse.Data = fmt.Sprintln("login error", lErr)
	} else {

		authenticatedUser = loginResponse.User
		cResponse.Data = fmt.Sprint("logged in successfully")
	}

}

func createTask(taskService *task.Service) {

	createRequest := cRequest.CreateTaskRequest

	if response, cErr := taskService.Create(task.CreateRequest{
		Title:               createRequest.Title,
		DueDate:             createRequest.DueDate,
		CategoryID:          createRequest.CategoryID,
		AuthenticatedUserID: authenticatedUser.ID,
	}); cErr != nil {
		cResponse.Data = fmt.Sprintln("task creation error", cErr)

		return
	} else {
		cResponse.Data = fmt.Sprintln("taskclient created :", response.Task)

	}
}
