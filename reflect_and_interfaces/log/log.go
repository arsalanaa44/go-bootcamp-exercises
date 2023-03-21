package log

import (
	"encoding/json"
	"os"
	"reflect_and_interfaces/anotherricherror"
	"reflect_and_interfaces/richerror"
	"time"
)

type Log struct {
	Errors []richerror.RichError
}

func (l *Log) Append(err error) {
	var richError *richerror.RichError
	if re, ok := err.(*richerror.RichError); ok {
		richError = re
	} else {
		if are, ok := err.(*anotherricherror.AnotherRichError); ok {
			richError = &richerror.RichError{
				Message:   are.Message,
				MetaData:  nil,
				Operation: are.Operation,
				Time:      time.Now(),
			}
		} else {
			richError = &richerror.RichError{
				Message:   err.Error(),
				MetaData:  nil,
				Operation: "unknown",
				Time:      time.Now(),
			}
		}
	}
	l.Errors = append(l.Errors, *richError)
}

func (l *Log) Save() {

	f, _ := os.OpenFile("errors.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	defer f.Close()

	jData, _ := json.Marshal(l.Errors)
	f.Write(jData)

}
