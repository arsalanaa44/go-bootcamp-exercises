package entity

type User struct {
	ID          int
	Name        string
	PhoneNumber string
	// Password keep hashed
	Password string
}
