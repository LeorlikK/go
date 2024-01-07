package cycle

import "fmt"

func For() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func While() {
	x := 0
	for {
		x++
		fmt.Println(x)
	}
}

func ForEach() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range slice {
		fmt.Println(slice[i])
	}
	for index, value := range slice {
		fmt.Println(index, value)
	}
	for _, value := range slice {
		fmt.Println(value)
	}
}

func Break() {
	counter := 0
	for {
		if counter == 100 {
			break
		}
		counter++
		fmt.Println(counter)
	}
}
