package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	//go get
	"golang.org/x/crypto/bcrypt"
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

var (
	userStorage       []User
	taskStorage       []Task
	categoryStorage   []Category
	authenticatedUser *User
	serializationMode string
)

const (
	userStoragePath               = "user.txt"
	MandaravardiSerializationMode = "mandaravardi"
	JsonSerializationMode         = "json"
)

func main() {

	loadUserStorageFromFile()
	for _, u := range userStorage {
		fmt.Printf("%+v\n", u)
	}

	fmt.Println("hello to TODO app")

	command := flag.String("command", "no command", "command to run")
	serializeMode := flag.String("serialization-mode", JsonSerializationMode, "serialization mode to write data to file")
	flag.Parse()

	switch *serializeMode {
	case MandaravardiSerializationMode:
		{
			serializationMode = MandaravardiSerializationMode
		}
	default:
		{
			serializationMode = JsonSerializationMode
		}
	}
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

	password = hashThePassword(password)
	fmt.Println(name, "registered")

	id := len(userStorage) + 1
	user := User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}
	userStorage = append(userStorage, user)

	writeUserToFile(user)

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

	for _, u := range userSlice {
		if u == "" {

			continue
		}
		var user User
		if serializationMode == MandaravardiSerializationMode {
			user, _ = deserializeFormMandaravardi(u)
		} else {
			var jErr error
			jErr = json.Unmarshal([]byte(u), &user)
			if jErr != nil {
				fmt.Println("error in Unmarshalization !", jErr)

				continue
			}
		}
		userStorage = append(userStorage, user)

	}
}
func writeUserToFile(user User) {

	file, err := os.OpenFile(userStoragePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open-file error :", err)

		return
	}
	defer file.Close()

	var data []byte
	if serializationMode == MandaravardiSerializationMode {
		data = []byte(fmt.Sprintf("\nID: %d, name: %s, email: %s, password: %s",
			user.ID, user.Name, user.Email, user.Password))
	} else if serializationMode == JsonSerializationMode {
		var er error
		data, er = json.Marshal(user)
		if er != nil {
			fmt.Println("can't marshal user to json", er)
		}
	}
	data = append(data, '\n')
	file.Write(data)
}
func deserializeFormMandaravardi(userStr string) (User, error) {
	var user User
	userStr = strings.ReplaceAll(userStr, " ", "")
	userFields := strings.Split(userStr, ",")
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
	return user, nil
}
func hashThePassword(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
