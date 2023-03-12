package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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

const userStoragePath = "user.txt"

func main() {

	loadUserStorageFromFile()
	for _, u := range userStorage {
		fmt.Printf("%+v\n", u)
	}

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
		fmt.Println("you should log in first !")
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
		fmt.Println("category-ID is not valid", e)
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

	writeToFile(user)

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

func loadUserStorageFromFile() {
	file, er := os.Open(userStoragePath)
	if er != nil {
		fmt.Println(userStoragePath, "doesn't exist")

		return
	}
	defer file.Close()

	data := make([]byte, 1024)
	file.Read(data)

	dataStr := string(data)
	userSlice := strings.Split(dataStr, "\n")

	for _, u := range userSlice[:len(userSlice)-1] {
		fmt.Println(len(u))
		var user User
		u = strings.ReplaceAll(u, " ", "")
		userFields := strings.Split(u, ",")
		for _, userField := range userFields {
			field := strings.Split(userField, ":")
			if len(field) < 2 {

				continue
			}
			fieldName := field[0]
			fieldValue := field[1]
			switch fieldName {
			case "ID":
				{
					var err error
					user.ID, err = strconv.Atoi(fieldValue)
					if err != nil {
						fmt.Println(err)
					}
				}
			case "name":
				{
					user.Name = fieldValue
				}
			case "password":
				{
					user.Password = fieldValue
				}
			case "email":
				{
					user.Email = fieldValue
				}
			default:
				{
					fmt.Println("hacker detected")
				}
			}
		}
		userStorage = append(userStorage, user)
	}

}
func writeToFile(user User) {
	file, err := os.OpenFile(userStoragePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open-file error :", err)

		return
	}
	defer file.Close()
	data := fmt.Sprintf("ID: %d, name: %s, email: %s, password: %s\n",
		user.ID, user.Name, user.Email, user.Password)
	file.Write([]byte(data))
}