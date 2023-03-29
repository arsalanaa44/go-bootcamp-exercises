package inmemory

import "todo_cli/entity"

type User struct {
	users []entity.User
}

func NewUserStore(users []entity.User) *User {
	return &User{
		users,
	}
}

func (u *User) Save(user entity.User) entity.User {

	user.ID = len(u.users) + 1
	u.users = append(u.users, user)

	return user
}

func (u *User) DeleteLast() {

	u.users = u.users[:len(u.users)-1]
}

func (u *User) ListUsers() []entity.User {
	var ls = make([]entity.User, len(u.users))
	copy(u.users, ls)

	return ls
}
