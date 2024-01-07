package str

import "fmt"

func String() {
	strOne := "STR_ONE"
	strTwo := "STR_TWO"
	intOne := 10

	fmt.Print("1")
	fmt.Printf("2 %s, %d", strOne, intOne)
	fmt.Println("3")
	res := fmt.Sprintf("string one, %s, plus string two: %s, and number - %d", strOne, strTwo, intOne)
	fmt.Println(res)
}
