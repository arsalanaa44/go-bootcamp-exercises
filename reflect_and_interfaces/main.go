package main

import (
	"fmt"
	"reflect_and_interfaces/app"
	"reflect_and_interfaces/mapmemory"
	"reflect_and_interfaces/user"
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

	application.CreateUser(user.User{
		ID:   1,
		Name: "Aa",
	})

	fmt.Println(application.ListUsers())
	fmt.Println(application.ListUsers())
	fmt.Println(application.GetUserByID(10))
	fmt.Println(application.GetUserByID(1))
}
