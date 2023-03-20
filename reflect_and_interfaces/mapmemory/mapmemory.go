package mapmemory

import "reflect_and_interfaces/user"

type MStore struct {
	Store map[int]user.User
}

func (m *MStore) CreateUser(user user.User) {
	m.Store[user.ID] = user
}

func (m *MStore) ListUser() []user.User {
	var userSlice []user.User
	for _, v := range m.Store {
		userSlice = append(userSlice, v)
	}

	return userSlice
}

func (m *MStore) GetUserByID(id int) user.User {

	return m.Store[id]
}
