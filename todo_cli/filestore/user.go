package filestore

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todo_cli/constant"
	"todo_cli/entity"
)

type FileStore struct {
	filePath          string
	serializationMode string
}

// constructor
func New(filePath, serializationMode string) FileStore {
	return FileStore{filePath, serializationMode}
}

func (fs FileStore) Save(user entity.User) {
	fs.writeUserToFile(user)
}

func (fs FileStore) writeUserToFile(user entity.User) {

	file, err := os.OpenFile(fs.filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open-file error :", err)

		return
	}
	defer file.Close()

	var data []byte
	if fs.serializationMode == constant.MandaravardiSerializationMode {
		data = []byte(fmt.Sprintf("\nID: %d, name: %s, email: %s, password: %s",
			user.ID, user.Name, user.Email, user.Password))
	} else if fs.serializationMode == constant.JsonSerializationMode {
		var er error
		data, er = json.Marshal(user)
		if er != nil {
			fmt.Println("can't marshal user to json", er)
		}
	}
	data = append(data, '\n')
	file.Write(data)
}

func (fs FileStore) Load() []entity.User {
	return fs.loadUserStorageFromFile()
}

func (fs FileStore) loadUserStorageFromFile() []entity.User {
	users := []entity.User{}
	file, er := os.Open(fs.filePath)
	if er != nil {
		fmt.Println(fs.filePath, "doesn't exist")

		return []entity.User{}
	}
	defer file.Close()

	data := make([]byte, 1024)
	file.Read(data)

	dataStr := string(data)
	userSlice := strings.Split(dataStr, "\n")

	for _, u := range userSlice {
		if !(u[0] == '{' || u[0] == 'I') {

			continue
		}
		var user entity.User
		if fs.serializationMode == constant.MandaravardiSerializationMode {
			user, _ = deserializeFormMandaravardi(u)
		} else {
			var jErr error
			jErr = json.Unmarshal([]byte(u), &user)
			if jErr != nil {
				fmt.Println("error in Unmarshalization !", jErr)

				continue
			}
		}
		users = append(users, user)

	}

	return users
}

func deserializeFormMandaravardi(userStr string) (entity.User, error) {
	var user entity.User
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
