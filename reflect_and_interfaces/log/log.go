package log

import (
	"encoding/json"
	"os"
	"reflect_and_interfaces/richerror"
)

type Log struct {
	Errors []richerror.RichError
}

func (l *Log) Append(richError richerror.RichError) {
	l.Errors = append(l.Errors, richError)
}

func (l *Log) Save() {

	f, _ := os.OpenFile("errors.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	defer f.Close()

	jData, _ := json.Marshal(l.Errors)
	f.Write(jData)

}
