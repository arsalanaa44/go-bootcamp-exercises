package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"reflect_and_interfaces/app"
	"reflect_and_interfaces/intj"
	"reflect_and_interfaces/log"
	"reflect_and_interfaces/mapmemory"
	"reflect_and_interfaces/richerror"
	"reflect_and_interfaces/simpledata"
	"reflect_and_interfaces/user"
	"sort"
	"time"
)

func main() {

	myMap := make(map[int]user.User)
	mapm := mapmemory.MStore{
		Store: myMap,
	}
	application := app.App{
		Name:        "myApp",
		UserStorage: &mapm,
	}

	var u = user.User{
		ID:   1,
		Name: "Aa",
	}
	application.CreateUser(u)

	application.CreateUser(user.User{
		ID:   3,
		Name: "Cc",
	})
	application.CreateUser(user.User{
		ID:   2,
		Name: "Bb",
	})

	fmt.Println(application.ListUsers())
	fmt.Println("see ?, map is not ordered :)")
	fmt.Println(application.GetUserByID(10))
	fmt.Println(application.GetUserByID(1))
	fmt.Println()

	var array = intj.Inta{1, 3, 5, 7, 6, 4, 2}
	fmt.Println(array)
	sort.Sort(array)
	fmt.Println("with Inta interface :\n", array)
	fmt.Println()

	var array0 = sort.IntSlice{1, 3, 5, 7, 6, 4, 2}
	fmt.Println(array0)
	sort.Sort(array0)
	fmt.Println("with IntSlice interface exactly the same as Inta interface:\n", array0)
	fmt.Println()

	var array2 = intj.Inta{1, 3, 5, 7, 6, 4, 2}
	fmt.Println(array2)
	sort.Ints(array2)
	fmt.Println("with Ints uses IntSlice interface :\n", array2)
	fmt.Println()

	//io.ReadWriter()
	//bufio.ReadWriter{}

	fmt.Println(u)
	if uM, jE := json.Marshal(u); jE == nil {
		fmt.Println(string(uM))
	}

	fmt.Println(user.CreateUserByID1(2))
	fmt.Println(user.CreateUserByID1(0))
	logger := log.Log{}

	filePath := "Im_not_exist.txt"
	if _, er := os.Open(filePath); er != nil {
		sEr := fmt.Sprint(er)
		r := &richerror.RichError{
			Message: "cannot open the file",
			MetaData: map[string]string{
				"error": sEr,
			},
			Operation: "os.open",
		}
		logger.Append(r)
		fmt.Println(1)
		fmt.Println(r.Error())
	}

	if _, er := user.CreateUserByID1(0); er != nil {

		logger.Append(er)
		fmt.Println(2)
		fmt.Println(er.Error())
	}

	if _, er := user.CreateUserByID2(0); er != nil {

		logger.Append(er)
		fmt.Println(3)
		fmt.Println(er.Error())
	}
	if _, er := user.CreateUserByID3(0); er != nil {

		logger.Append(er)
		fmt.Println(3)
		fmt.Println(er.Error())
	}

	richError := richerror.RichError{
		Message:   "a simple richError",
		MetaData:  nil,
		Operation: "operand",
		Time:      time.Now(),
	}
	value := reflect.ValueOf(richError)
	switch value.Kind() {
	case reflect.Struct:
		{
			for i := 0; i < value.NumField(); i++ {
				fmt.Println(i,
					value.Field(i),
					value.Type(),
					value.Type().Field(i).Name,
					value.Type().Field(i).Type)
			}
		}
	}

	logger.Save()

}

func prt(err error) {
	fmt.Sprintln(err)
	fmt.Sprintln(json.Marshal(simpledata.SimpleDataTwo{
		ID:    00,
		Name:  "name",
		Email: "email",
	}))

}

func prtDotErr(err error) {
	fmt.Sprintln(err.Error())
	fmt.Sprintln(json.Marshal(simpledata.SimpleData{
		ID:    00,
		Name:  "name",
		Email: "email",
	}))
}
