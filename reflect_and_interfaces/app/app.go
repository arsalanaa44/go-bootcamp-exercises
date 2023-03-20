package app

import "reflect_and_interfaces/user"

type UserStore interface {
	CreateUser(user user.User)
	ListUser() []user.User
	GetUserByID(int) user.User
}
type App struct {
	Name        string
	UserStorage UserStore
}

func (a App) CreateUser(user user.User) {
	a.UserStorage.CreateUser(user)
}

func (a App) ListUsers() []user.User {

	return a.UserStorage.ListUser()
}

func (a App) GetUserByID(id int) user.User {

	return a.UserStorage.GetUserByID(id)
}
