package richerror

import (
	"fmt"
	"time"
)

// error is an interface
type RichError struct {
	Message   string
	MetaData  map[string]string
	Operation string
	Time      time.Time
}

func (r *RichError) Error() string {

	return r.Message
}

func (r *RichError) String() string {
	return fmt.Sprintf("\nMessage:%s\nMetaData:%+v\nOperation:%s\nTime:%v",
		r.Message,
		r.MetaData,
		r.Operation,
		r.Time)
}
