package array

import "fmt"

func Array() {
	// Копия
	test := [3]int{
		1, 2, 3,
	}
	var testTwo [10]string

	test[0] = 11
	testFunc := func(test [3]int) {
		test[0] = 10
	}
	fmt.Println(test)
	fmt.Println(test[0])
	// testTwo[0] = "1"
	fmt.Println(testTwo)
	testFunc(test)
	fmt.Println(test)
}

func Slice() {
	// Передача по ссылке
	a := []int{1, 2, 3, 4, 5}
	i := 2
	a[i] = a[len(a)-1]
	a = a[:len(a)-1]
	fmt.Println(a)

	test := []int{1, 2, 3, 4, 5}
	var testTwo []string
	testThree := make([]int, 5, 5)
	test[i] = test[len(test)-1]
	fmt.Println(test)
	testFunc := func(test []int) {
		test[0] = 10
	}
	// testTwo[0] = "1"
	testThree[0] = 1
	fmt.Println(testTwo)
	fmt.Println(testThree, len(testThree), cap(testThree))
	testThree = append(testThree, 6, 7, 8)
	fmt.Println(testThree, len(testThree), cap(testThree))
	testFunc(test)
	fmt.Println(test)
}

func Map() {
	// Map
	test := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	// var testTwo map[string]string
	testThree := make(map[string]bool)

	fmt.Println(test)
	fmt.Println(test["one"])
	fmt.Println(test["four"])
	four, ex := test["four"]
	if ex {
		fmt.Println(four)
	} else {
		fmt.Println("Not founded")
	}
	test["four"] = 4
	delete(test, "one")
	for key, value := range test {
		fmt.Println(key, value)
	}

	// testTwo["one"] = "1"

	testThree["one"] = true
	testThree["two"] = false
	testThree["three"] = true

	fmt.Println(testThree)
	fmt.Println(len(testThree))
}

func Matrix() {
	// matrix := [][]int{}
	matrix := make([][]int, 10)

	for x := 0; x < 10; x++ {
		matrix[x] = make([]int, 10)
		matrix[x][x] = x
		fmt.Println(matrix)
	}
}
