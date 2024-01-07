package interfaceTest

type UserInteface interface {
	Sum() int
	FuncReturnInterface() interface{}
	FuncReturnNotNullInterface() ReturnInterface
}

type ReturnInterface interface {
	ConcatString() string
}

type UserOne struct {
	X int
	Y int
}

type UserTwo struct {
	A int
	B int
}

func (u UserOne) Sum() int {
	return u.X + u.Y
}
func (u UserOne) FuncReturnInterface() interface{} {
	// valueString := "string"
	// valueIntager := 50
	// valueBool := false
	valueSlice := make([]int, 0)
	return valueSlice
}
func (u UserOne) FuncReturnNotNullInterface() ReturnInterface {
	return UserOne{
		X: 10,
		Y: 20,
	}
	// return UserTwo{
	// A: 10,
	// B: 20,
	// }
}
func (u UserOne) ConcatString() string {
	return "string"
}

func (u UserTwo) Sum() int {
	return u.A + u.B
}
func (u UserTwo) FuncReturnInterface() interface{} {
	valueString := "string"
	return valueString
}

func NewOneUser(X, Y int) UserOne {
	return UserOne{X, Y}
}
