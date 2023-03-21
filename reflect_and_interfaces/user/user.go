package user

import (
	"fmt"
	"reflect_and_interfaces/anotherricherror"
	"reflect_and_interfaces/richerror"
	"strconv"
	"time"
)

type User struct {
	ID   int
	Name string
}

// implement stringer-interface methods
func (u User) String() string {
	return fmt.Sprintf("---\nID : %d\nName : %s\n---\n", u.ID, u.Name)
}

func CreateUserByID1(id int) (User, error) {
	if id == 0 {
		r := richerror.RichError{
			Message:   "id is 0",
			MetaData:  nil,
			Operation: "CreateUserByID",
			Time:      time.Now(),
		}

		return User{}, &r

	}

	return User{ID: id, Name: ("number" + strconv.Itoa(id))}, nil
}

func CreateUserByID2(id int) (User, error) {
	if id == 0 {
		r := anotherricherror.AnotherRichError{
			Message:   "id is 0",
			Operation: "CreateUserByID2",
		}

		return User{}, &r

	}

	return User{ID: id, Name: ("number" + strconv.Itoa(id))}, nil
}

// can't assert to richError or anotherRichError
func CreateUserByID3(id int) (User, error) {
	if id == 0 {

		return User{}, fmt.Errorf("id is 0")
	}

	return User{ID: id, Name: ("number" + strconv.Itoa(id))}, nil
}
