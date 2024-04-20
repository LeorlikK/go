package error

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var (
	muErr sync.Mutex

	Error404 = NewCustomError("404M", "404C", true, nil)
)

type CustomError struct {
	Message string
	Code    string
	Log     bool
	Err     error
}

func (e *CustomError) Error() string {
	logger := e.ErrorString()

	if e.Log {
		e.ErrorLog(logger)
	}

	return logger
}

func (e *CustomError) Unwrap() error {
	return e.Err
}

func NewCustomError(message, code string, log bool, err error) error {
	return &CustomError{
		Message: message,
		Code:    code,
		Log:     log,
		Err:     err,
	}
}

func (e *CustomError) ErrorLog(logger string) {
	muErr.Lock()
	defer muErr.Unlock()

	file, err := os.OpenFile("./error/err.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	if _, err = file.WriteString(logger + "\n"); err != nil {
		log.Fatal(err)
	}
}

func (e *CustomError) ErrorString() string {
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	return fmt.Sprintf("%s Message: %s, Code: %s, Err: %s", currentTime, e.Message, e.Code, e.Err)
}
