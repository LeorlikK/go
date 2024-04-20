package error

import (
	"fmt"
	"time"
)

type MyErrorStruct struct {
}

func (e *MyErrorStruct) Error() string {
	func() {
		fmt.Println("log_1.log_1.log_1")
	}()
	return "my error struct"
}

type MyErrorStructTwo struct {
}

func (e *MyErrorStructTwo) Error() (string, error) {
	return "", nil
}

var messageError = "error one"

func Error() (string, error) {
	//err := errors.New("new error")
	//err = fmt.Errorf("new error[%s]", messageError)

	return "my error struct Two", NewCustomError("111", "222", false, nil)
}

func ErrorTest() string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("RECOVER")
		}

		fmt.Println("RECOVER")
	}()

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		go ErrorGTwo()
		time.Sleep(1 * time.Second)
	}
	fmt.Println("END MAIN")
	return "str"
}

func ErrorGTwo() {
	arr := []int{1, 2, 3}
	fmt.Println(arr[10])
}
