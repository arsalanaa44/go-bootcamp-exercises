package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}
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

var userStorage []User
var taskStorage []Task
var categoryStorage []Category
var authenticatedUser *User

func main() {

	fmt.Println("hello to TODO app")
	command := flag.String("command", "no command", "command to run")
	flag.Parse()

	for {
		runCommand(*command)

		fmt.Println("please enter another command :")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		*command = scanner.Text()

	}

}

func runCommand(command string) {
	if command != "register-user" && command != "login" && command != "exit" && authenticatedUser == nil {
		login()
		if authenticatedUser == nil {

			return
		}
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
	case "create-category":
		{
			createCategory()
		}
	case "register-user":
		{
			registerUser()
		}
	case "login":
		{

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
	if e == nil {
		fmt.Println("category-ID is not valid integer", e)
	}

	notFound := true
	for _, c := range categoryStorage {
		if c.ID == categoryID && c.userID == authenticatedUser.ID {
			notFound = false

			break
		}
	}
	if notFound {
		fmt.Println("category-ID is not valid", e)
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

func registerUser() {
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

	fmt.Println("registered", email, password)

	id := len(userStorage) + 1
	user := User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}
	userStorage = append(userStorage, user)

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
		if user.Password == password && user.Email == email {
			authenticatedUser = &user
			fmt.Println("user login :", email, password)

			break
		}
	}
	if authenticatedUser == nil {
		fmt.Println("email or password is in correct")
	}

}

func listTask() {

	for _, task := range taskStorage {
		if task.UserID == authenticatedUser.ID {
			fmt.Printf("%+v\n", task)
		}
	}
}
