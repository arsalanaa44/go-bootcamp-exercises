package main

import (
	"encoding/json"
	"fmt"
	"reflect_and_interfaces/app"
	"reflect_and_interfaces/intj"
	"reflect_and_interfaces/mapmemory"
	"reflect_and_interfaces/user"
	"sort"
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

}
