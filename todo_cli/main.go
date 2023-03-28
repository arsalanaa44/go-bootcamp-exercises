package main

import (
	"bufio"
	"flag"
	"fmt"
	"todo_cli/constant"
	"todo_cli/contract"
	"todo_cli/encryption"
	"todo_cli/entity"
	"todo_cli/repository/filestore"
	"todo_cli/repository/inmemory"
	"todo_cli/service/category"
	"todo_cli/service/task"

	//go get
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
)

var (
	userStorage       []entity.User
	categoryStorage   []entity.Category
	authenticatedUser *entity.User
	serializationMode string
)

const (
	userStoragePath = "user.txt"
)

func main() {

	taskMemoryRepo := inmemory.NewTaskStore()
	categoryMemoryRepo := inmemory.NewCategoryStore()
	categoryService := category.NewService(categoryMemoryRepo)
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

	//loadUserStorageFromFile()

	//var store userReadStore
	//var userFileStore = fileStore{userStoragePath}
	//store = userFileStore
	//loadUserStorage(store)
	userStorage = userFileStore.Load()

	for {
		runCommand(userFileStore, *command, &taskService)

		fmt.Println("please enter another command :")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		*command = scanner.Text()

	}

}

func runCommand(store contract.UserWriteStore, command string, taskService *task.Service) {
	if command != "register-user" && command != "login" && command != "exit" && authenticatedUser == nil {
		fmt.Println("you should log in first !")
		login()
		if authenticatedUser == nil {

			return
		}
	}
	if command == "register-user" {
		//var store userWriteStore
		//var userFileStore = fileStore{userStoragePath}
		//store = userFileStore
		registerUser(store)

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
			listUser()
		}
	case "create-category":
		{
			createCategory()
		}
	case "login":
		{
			login()
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
	scanner := bufio.NewScanner(os.Stdin)
	var title, dueDate, category string

	fmt.Println("please enter the task title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the task due date")
	scanner.Scan()
	dueDate = scanner.Text()

	fmt.Println("please enter the task category-ID")
	scanner.Scan()
	category = scanner.Text()
	categoryID, e := strconv.Atoi(category)
	if e != nil {
		fmt.Println("category-ID is not valid integer", e)

		return
	}

	if response, cErr := taskService.Create(task.CreateRequest{
		Title:               title,
		DueDate:             dueDate,
		CategoryID:          categoryID,
		AuthenticatedUserID: authenticatedUser.ID,
	}); cErr != nil {
		fmt.Println("error", cErr)

		return
	} else {
		fmt.Println("task created :", response.Task)

	}

}

func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)
	var title, color string

	fmt.Println("please enter the category title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the category color")
	scanner.Scan()
	color = scanner.Text()

	fmt.Println("category: ", title, color)

	category := entity.Category{
		ID:     len(categoryStorage) + 1,
		Title:  title,
		Color:  color,
		UserID: authenticatedUser.ID,
	}
	categoryStorage = append(categoryStorage, category)

}

func registerUser(store contract.UserWriteStore) {

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
	fmt.Println(name, "registered")

	id := len(userStorage) + 1
	user := entity.User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}
	userStorage = append(userStorage, user)

	// writeUserToFile(user)
	store.Save(user)
}

func login() {

	fmt.Println("login process :")
	scanner := bufio.NewScanner(os.Stdin)
	var email, password string

	fmt.Println("please enter the user email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the user password")
	scanner.Scan()
	password = scanner.Text()
	for _, user := range userStorage {
		if user.Email == email {
			er := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if er == nil {
				authenticatedUser = &user
				fmt.Println("user login :", user.Name)

				break
			}
		}
	}
	if authenticatedUser == nil {
		fmt.Println("email or password is incorrect")
	}

}

func listTask(taskService *task.Service) {

	lReq := task.ListRequest{authenticatedUser.ID}
	if listResponse, lErr := taskService.List(lReq); lErr != nil {
		fmt.Println("error listTask", lErr)

		return
	} else {
		for i, v := range listResponse.Tasks {
			fmt.Println(i+1, ":", v)
		}
	}

	return
}

func listUser() {
	for _, u := range userStorage {
		fmt.Printf("%+v\n", u)
	}
}


