package richerror

import "fmt"

// error is an interface
type RichError struct {
	Message   string
	MetaData  map[string]string
	Operation string
}

func (r *RichError) Error() string {

	return r.String()
}

func (r RichError) String() string {
	return fmt.Sprintf("\nMessage:%s\nMetaData:%+v\nOperation:%s",
		r.Message,
		r.MetaData,
		r.Operation)
}
