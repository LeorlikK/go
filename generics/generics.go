package generics

import "fmt"

func Sum[T int64 | float64](slice []T) (sum T) {
	sum = 0
	for _, value := range slice {
		sum += value
	}
	return
}

type Numeric interface {
	int64 | float64
}

type Numbers[T Numeric] []T

var ints Numbers[int64]

func Contains[T Numeric](slice []T, element T) (bool, T) {
	for _, el := range slice {
		if el == element {
			return true, el
		}
	}

	return false, 0
}

func Contains2[T comparable](slice []T, element T) bool {
	for _, el := range slice {
		if el == element {
			return true
		}
	}

	return false
}

func Contains3[T any](elements ...T) {
	fmt.Println(elements)
}
