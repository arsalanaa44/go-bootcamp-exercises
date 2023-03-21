package simpledata

import "fmt"

type SimpleData struct {
	ID    int
	Name  string
	Email string
}

func (s SimpleData) MarshalJSON() ([]byte, error) {
	return []byte(s.String()), nil
}
func (s SimpleData) String() string {
	return fmt.Sprintf("{\"ID\":\"%d\",\"Name\":\"%s\",\"Email\":\"%s\"}",
		s.ID,
		s.Name,
		s.Email)
}
