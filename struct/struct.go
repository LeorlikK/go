package structTest

import (
	"fmt"
)

type OneStruct struct {
	Name    string
	Family  string
	Age     int
	secret  int //private
	static  int
	Array   [3]int
	Slice   []string
	Mapping map[string]string
}

func NewOneStruct(
	name, family string,
	age, secret int,
	array [3]int,
	slice []string,
	mapping map[string]string,
) OneStruct {
	return OneStruct{
		Name:    name,
		Family:  family,
		Age:     age,
		secret:  secret,
		static:  12345,
		Array:   array,
		Slice:   slice,
		Mapping: mapping,
	}
}

func (o OneStruct) PrintObj() {
	fmt.Println(o)
}

func (o OneStruct) GetValue(key string) interface{} {
	switch key {
	case "Name":
		return o.Name
	case "Family":
		return o.Family
	case "Age":
		return o.Age
	case "secret":
		return o.secret
	case "static":
		return o.static
	default:
		return nil
	}
}

func (o *OneStruct) SetValue(key string, value interface{}) {
	switch key {
	case "Name":
		if v, ok := value.(string); ok {
			o.Name = v
		}
	case "Family":
		if v, ok := value.(string); ok {
			o.Family = v
		}
	case "Age":
		if v, ok := value.(int); ok {
			o.Age = v
		}
	case "secret":
		if v, ok := value.(int); ok {
			o.secret = v
		}
	case "static":
		if v, ok := value.(int); ok {
			o.static = v
		}
	}
}
