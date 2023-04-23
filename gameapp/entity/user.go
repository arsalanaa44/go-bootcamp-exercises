package entity

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	// Password keep hashed
	Password string // `json:"-"` // secure in marshalization but not good solution
}
