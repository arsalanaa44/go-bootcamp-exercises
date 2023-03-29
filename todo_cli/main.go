package main

import (
	"bufio"
	"flag"
	"fmt"
	"todo_cli/constant"
	"todo_cli/encryption"
	"todo_cli/entity"
	"todo_cli/repository/filestore"
	"todo_cli/repository/inmemory"
	"todo_cli/service/categoryservice"
	"todo_cli/service/task"
	"todo_cli/service/user"

	"os"
)

var (
	authenticatedUser *entity.User
	serializationMode string
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

	command := flag.String("command", "no command", "command to run")
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

	//loadUserStorageFromFile()

	//var store userReadStore
	//var userFileStore = fileStore{userStoragePath}
	//store = userFileStore
	//loadUserStorage(store)

	for {
		runCommand(&userService, *command, &taskService, &categoryService)

		fmt.Println("please enter another command :")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		*command = scanner.Text()

	}

}

func runCommand(userService *user.Service, command string, taskService *task.Service, categoryService *categoryservice.Service) {
	if command != "register-user" && command != "login" && command != "exit" && authenticatedUser == nil {
		fmt.Println("you should log in first !")
		login(userService)
		if authenticatedUser == nil {

			return
		}
	}
	if command == "register-user" {
		//var store userWriteStore
		//var userFileStore = fileStore{userStoragePath}
		//store = userFileStore
		registerUser(userService)

		return
	}

	switch command {

	case "create-task":
		{
			createTask(taskService)
		}
	case "list-task":
		{
			listTask(taskService)
		}
	case "list-user":
		{
			listUser(userService)
		}
	case "create-category":
		{
			createCategory(categoryService)
		}
	case "login":
		{
			login(userService)
		}
	case "exit":
		{
			os.Exit(0)
		}
	default:
		{
			fmt.Println("command is not valid !")
		}
	}
}

func createTask(taskService *task.Service) {

	//createRequest, cErr := taskclient.NewClientService().CreateTask()
	//if cErr != nil {
	//	fmt.Println("can't create task", cErr)
	//
	//	return
	//}
	if response, cErr := taskService.Create(task.CreateRequest{
		Title:               "createRequest.Title",
		DueDate:             "createRequest.DueDate",
		CategoryID:          1, //createRequest.CategoryID,
		AuthenticatedUserID: authenticatedUser.ID,
	}); cErr != nil {
		fmt.Println("error", cErr)

		return
	} else {
		fmt.Println("taskclient created :", response.Task)

	}

}

func createCategory(categoryService *categoryservice.Service) {
	//scanner := bufio.NewScanner(os.Stdin)
	//var title, color string
	//
	//fmt.Println("please enter the categoryservice title")
	//scanner.Scan()
	//title = scanner.Text()
	//
	//fmt.Println("please enter the categoryservice color")
	//scanner.Scan()
	//color = scanner.Text()

	createResponse, cErr := categoryService.Create(categoryservice.CreateRequest{
		Title:               "title",
		Color:               "color",
		AuthenticatedUserID: authenticatedUser.ID,
	})
	if cErr != nil {
		fmt.Println("categoriy creation error:", cErr)

		return
	} else {
		fmt.Println("category created:", createResponse.Category.Title)
	}

}

func registerUser(userService *user.Service) {

	scanner := bufio.NewScanner(os.Stdin)
	var name, email, password string

	fmt.Println("please enter the user name")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the user email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the user password")
	scanner.Scan()
	password = scanner.Text()

	password = encryption.HashThePassword(password)

	registerRes, rErr := userService.Register(user.RegisterRequest{
		name,
		email,
		password,
	})
	if rErr != nil {
		fmt.Println("error in registering user", rErr)
	}
	fmt.Println(registerRes.User.Name, "registered")

}

func login(userService *user.Service) {
	//loginRequest := userclient.NewClientService().Login()
	//autologin
	loginRequest := user.LoginRequest{
		Email:    "pp",
		Password: "ppp",
	}
	if loginResponse, lErr := userService.Login(loginRequest); lErr != nil {
	} else {
		authenticatedUser = loginResponse.User
		fmt.Println(authenticatedUser.Name, "logged in successfully")
	}

}

func listTask(taskService *task.Service) {

	lReq := task.ListRequest{authenticatedUser.ID}
	if listResponse, lErr := taskService.List(lReq); lErr != nil {
		fmt.Println("error listTask", lErr)

		return
	} else {
		println("what the")
		for i, v := range listResponse.Tasks {
			fmt.Println(i+1, ":", v)
		}
	}

	return
}

func listUser(userService *user.Service) {
	for _, u := range userService.List().Users {
		fmt.Printf("%+v\n", u)
	}
}
