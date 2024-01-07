package arguments

import (
	"fmt"
	"reflect"
)

func Main() {
	min := FindMinNumber(2, 1, 4, 3, 6, 5, 8, 7)
	text, intager, max := FindMaxNumber("one", 111, 2, 1, 4, 3, 6, 5, 8, 7)
	fmt.Println(min)
	fmt.Println(text, intager, max)
}

func FindMinNumber(numbers ...int) int {
	fmt.Println(reflect.TypeOf(numbers))
	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]

	for _, value := range numbers {
		if value < min {
			min = value
		}
	}

	return min
}

func FindMaxNumber(text string, intager int, kwarg ...int) (string, int, []int) {
	return text, intager, kwarg
}
