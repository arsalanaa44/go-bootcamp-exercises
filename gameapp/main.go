package main

import (
	"fmt"
	"gameapp/entity"
	"gameapp/repository/mysql"
)

func main() {

	mysglRepo := mysql.New()
	if createdUser, rErr := mysglRepo.Register(entity.User{
		Name:        "user",
		PhoneNumber: "122309483",
	}); rErr != nil {
		fmt.Printf("error in registering user :%w \n", rErr)
	} else {
		fmt.Println(createdUser)
	}
	fmt.Println(mysglRepo.IsPhoneNumberUnique("239232320"))
	fmt.Println(mysglRepo.IsPhoneNumberUnique("122309483"))

}
