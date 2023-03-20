package user

import "fmt"

type User struct {
	ID   int
	Name string
}

// implement stringer-interface methods
func (u User) String() string {
	return fmt.Sprintf("---\nID : %d\nName : %s\n---\n", u.ID, u.Name)
}
