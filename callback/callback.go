package callback

import "fmt"

func Main() {
	callback := func(str string, intager int) string {
		join := fmt.Sprintf("%s + %d end", str, intager)

		return join
	}

	res := callback("one", 1)
	fmt.Println(res)

	callbackFunc := Callback()
	resultCallback := callbackFunc(1)
	fmt.Println(resultCallback)
	resultCallback = callbackFunc(1)
	fmt.Println(resultCallback)
}

func Callback() func(integer int) int {
	count := 0

	return func(integer int) int {
		// count := 0
		count += integer
		return count
	}
}

//PHP
// function callback() {
//     $count = 0;

//     return function (int $integer) use (&$count): int {
//         $count += $integer;
//         return $count;
//     };
// }
