package main

import (
	"fmt"
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
	// test.SyncTest()

	// __GOROUTINE__
	//goroutines.Main()

	// __SOCKET__
	//socket.Server()
	//gorilla.Server()

	// __GENERICS__
	//int64Slice := []int64{20, 30, 40, 50}
	//sumInt64 := generics.Sum[int64](int64Slice)
	//fmt.Println(sumInt64)
	//
	//float64Slice := []float64{20.5, 30.3, 40.2, 50.1}
	//sumFloat64 := generics.Sum[float64](float64Slice)
	//fmt.Println(sumFloat64)
	//
	//numericSlice := []int64{1, 2, 3, 4, 5}
	//find, _ := generics.Contains(numericSlice, 8)
	//fmt.Println(find)

	// __ERROR__
	//if err := myError.Error(); err != nil {
	//	fmt.Println(err)
	//}

	// __REQUEST__
	//request.Request()

	// __SYNC__
	//syncMy.WithSlice()
	//syncMy.WithMap()

	// __PROJECTION__
	//projection.Main()

	// __CONVERT_TYPE__
	//convert_type.Main()

	// __ASSERT_TYPE__
	//assert_type.Main()

	// __DATABASE__
	//database.Main()

	// __CTX__
	//ctx.Main()
}

type IStruct interface {
	changeName(newName string, printer IStruct)
	printName()
	changeNameAfter()
}
type One struct {
	Name string
}

func (o *One) changeName(newName string, printer IStruct) {
	o.Name = newName
	printer.printName()
	printer.changeNameAfter()
}

func (o *One) printName() {
	fmt.Println(o.Name)
}
func (o *One) changeNameAfter() {
	fmt.Println("changeNameAfter")
}

type Two struct {
	One
}

func (o *Two) changeNameAfter() {
	fmt.Println("Before: ", o.Name)
	o.Name = fmt.Sprintf("Change_%s", o.Name)
	fmt.Println(o.Name)
}
func (o *Two) printName() {
	fmt.Println("Ничего я тебе не принтану!")
}
