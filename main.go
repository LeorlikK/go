package main

import (
	// "education/callback"
	// "education/arguments"
	// "education/string"
	// "education/link"
	// "education/cycle"
	// "education/array"
	// "education/defer"
	// "education/defer"
	// "education/struct"
	// "education/interface"
	// "education/test"
	// "education/goroutines"
	"education/goroutines"
	"fmt"
	// "testing"
	// "reflect"
	// "reflect"
)

var startMessage string

func init() {
	startMessage = "START"
	fmt.Println(startMessage)
}

func main() {
	// __STRING__
	// str.String()

	// __LINK__
	// link.Link()

	// __CYCLE__
	// cycle.For()
	// cycle.While()
	// cycle.ForEach()
	// cycle.Break()

	// __ARRAY__
	// array.Array()
	// array.Slice()
	// array.Map()
	// array.Matrix()

	// __ARGIMENTS__
	// arguments.Main()

	// __CALLBACK__
	// callback.Main()

	// __DEFER__
	// deferPackage.Main()

	// __STRUCT__
	// objOne := structTest.OneStruct{"Kirill", "Orlov", 27, 321}
	// objTwo := structTest.NewOneStruct(
	// 	"Masha",
	// 	"Kets",
	// 	28,
	// 	123,
	// 	[3]int{111, 222, 333},
	// 	make([]string, 0),
	// 	make(map[string]string),
	// )
	// objTwo.PrintObj()
	// family := objTwo.GetValue("Family")
	// age := objTwo.GetValue("Age")
	// fmt.Println(reflect.TypeOf(family), reflect.TypeOf(age))
	// objTwo.Array[0] = 444
	// objTwo.Slice = append(objTwo.Slice, "one", "two", "three")
	// objTwo.Slice = objTwo.Slice[1:]
	// objTwo.Slice[1] = objTwo.Slice[len(objTwo.Slice)-1]
	// objTwo.Slice = objTwo.Slice[:len(objTwo.Slice)-1]
	// objTwo.Mapping["one"] = "one_string"
	// objTwo.Mapping["two"] = "two_string"
	// delete(objTwo.Mapping, "two")
	// objTwo.PrintObj()

	// __INTERFACE__
	// var interfaceOne interfaceTest.InterfaseTest
	// interfaceOne = interfaceTest.StructTwo{1, 2}

	// structOne := interfaceTest.NewOneUser(1, 2)
	// structTwo := interfaceTest.NewOneUser(3, 4)

	// funcOne := func(u interfaceTest.UserInteface) int {
	// 	sum := u.Sum()
	// 	return sum
	// }
	// funcReturnInterface := func(u interfaceTest.UserInteface) interface{} {
	// 	return u.FuncReturnInterface()
	// }
	// FuncReturnNotNullInterface := func(u interfaceTest.UserInteface) interfaceTest.ReturnInterface {
	// 	return u.FuncReturnNotNullInterface().(interfaceTest.ReturnInterface) // разобраться позже!
	// }
	// sumOne := funcOne(structTwo)
	// resultReturnInterface := funcReturnInterface(structTwo)
	// resultFuncReturnNotNullInterface := FuncReturnNotNullInterface(structTwo)

	// fmt.Println(sumOne)
	// fmt.Println(resultReturnInterface, reflect.TypeOf(resultReturnInterface))
	// fmt.Println(resultFuncReturnNotNullInterface, reflect.TypeOf(resultFuncReturnNotNullInterface))

	// var a interface{}
	// a = "string"
	// fmt.Printf("%v, %T\n", a, a)
	// a = 42
	// fmt.Printf("%v, %T\n", a, a)
	// a = "42"
	// fmt.Printf("%v, %T\n", a, a)

	// __TESTING__
	// t := new(testing.T)
	// test.TestMultiple(t)

	// __GOROUTINE__
	goroutines.Main()

	fmt.Println("END")
}
