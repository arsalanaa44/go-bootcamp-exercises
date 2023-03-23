package main

import (
	"bufio"
	"flag"
	"fmt"
	"todo_cli/constant"
	"todo_cli/contract"
	"todo_cli/entity"

	//go get
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"todo_cli/filestore"
)

type Task struct {
	ID         int
	Title      string
	DueDate    string
	CategoryID int
	IsDone     bool
	UserID     int
}
type Category struct {
	ID     int
	Title  string
	Color  string
	userID int
}

var (
	userStorage       []entity.User
	taskStorage       []Task
	categoryStorage   []Category
	authenticatedUser *entity.User
	serializationMode string
)

const (
	userStoragePath = "user.txt"
)

func main() {

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
		runCommand(userFileStore, *command)

		fmt.Println("please enter another command :")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		*command = scanner.Text()

	}

}

func runCommand(store contract.UserWriteStore, command string) {
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
			createTask()
		}
	case "list-task":
		{
			listTask()
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

func createTask() {
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

	notFound := true
	for _, c := range categoryStorage {
		if c.ID == categoryID && c.userID == authenticatedUser.ID {
			notFound = false

			break
		}
	}
	if notFound {
		fmt.Println("category-ID is not valid")
		return
	}
	task := Task{
		ID:         len(taskStorage) + 1,
		Title:      title,
		DueDate:    dueDate,
		CategoryID: categoryID,
		IsDone:     false,
		UserID:     authenticatedUser.ID,
	}

	taskStorage = append(taskStorage, task)

	fmt.Println("task", title, dueDate, category)

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

	category := Category{
		ID:     len(categoryStorage) + 1,
		Title:  title,
		Color:  color,
		userID: authenticatedUser.ID,
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

	password = hashThePassword(password)
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

func listTask() {

	for _, task := range taskStorage {
		if task.UserID == authenticatedUser.ID {
			fmt.Printf("%+v\n", task)
		}
	}
}

func listUser() {
	for _, u := range userStorage {
		fmt.Printf("%+v\n", u)
	}
}

func hashThePassword(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
