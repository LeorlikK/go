package deferPackage

import "fmt"

func Main() {
	defer handlerDefer()

	array := []int{1, 2, 3, 4, 5}
	fmt.Println(array[1])
	fmt.Println(array[10])
	fmt.Println("FINISH")
}

func handlerDefer() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}

	fmt.Println("Success!")
}
