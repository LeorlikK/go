package projection

import "fmt"

type MyInterface interface {
	Method1() string
	Method2() string
	Method3() string
}

type MyBaseStruct struct{ Name string }

func (s MyBaseStruct) Method1() string {
	return "Method1 from MyBaseStruct"
}
func (s MyBaseStruct) Method2() string {
	return "Method2 from MyBaseStruct"
}

type MyCustomStruct struct {
	MyBaseStruct
}
type TestStruct struct{}

func (s MyCustomStruct) Method2() string {
	return "Method2 from MyCustomStruct"
}

func (s MyCustomStruct) Method3() string {
	return "Method3 from MyCustomStruct and field Name: " + s.Name
}

func NewMyCustomStruct() MyInterface {
	myCustomStruct := &MyCustomStruct{}
	myCustomStruct.Name = "LALA"
	return myCustomStruct
}

func Main() {
	base := MyBaseStruct{}
	fmt.Println(base.Method1())
	fmt.Println(base.Method2())

	custom := MyCustomStruct{}
	customCreate := NewMyCustomStruct()
	fmt.Println(custom.Method1())
	fmt.Println(custom.Method2())
	fmt.Println(custom.Name)

	//func(obj MyInterface) {
	//	fmt.Println(obj.Method1())
	//	fmt.Println(obj.Method2())
	//}(base)

	func(obj MyInterface) {
		fmt.Println(obj.Method1())
		fmt.Println(obj.Method2())
		fmt.Println(obj.Method3())
	}(custom)

	func(obj MyInterface) {
		fmt.Println(obj.Method1())
		fmt.Println(obj.Method2())
		fmt.Println(obj.Method3())
	}(customCreate)
}
