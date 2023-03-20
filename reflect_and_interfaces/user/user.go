package user

import (
	"fmt"
	"reflect_and_interfaces/richerror"
	"strconv"
)

type User struct {
	ID   int
	Name string
}

// implement stringer-interface methods
func (u User) String() string {
	return fmt.Sprintf("---\nID : %d\nName : %s\n---\n", u.ID, u.Name)
}

func CreateUserByID(id int) (User, error) {
	if id == 0 {
		r := richerror.RichError{
			Message:   "id is 0",
			MetaData:  nil,
			Operation: "CreateUserByID",
		}

		return User{}, &r

	}

	return User{ID: id, Name: ("number" + strconv.Itoa(id))}, nil
}
