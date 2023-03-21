package anotherricherror

import "fmt"

// error is an interface
type AnotherRichError struct {
	Message   string
	Operation string
}

func (r *AnotherRichError) Error() string {

	return r.Message
}

func (r *AnotherRichError) String() string {
	return fmt.Sprintf("\nMessage:%s\nOperation:%s",
		r.Message,
		r.Operation)
}
