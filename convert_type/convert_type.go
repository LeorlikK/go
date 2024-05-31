package convert_type

import (
	"fmt"
	"strconv"
)

func Main() {
	// String
	baseString := "15"
	intString, _ := strconv.Atoi(baseString)
	fmt.Println(intString)

	baseString = "3.1423423"
	float64String, _ := strconv.ParseFloat(baseString, 64)
	fmt.Println(float64String)

	baseString = "true"
	boolString, _ := strconv.ParseBool(baseString)
	fmt.Println(boolString)

	baseString = "random_string"
	byteString := []byte(baseString)
	fmt.Println(byteString)

	// Int
	baseInt := 15
	stringInt := fmt.Sprintf("%d", baseInt)
	fmt.Println(stringInt)

	baseInt = 15
	float64Int := float64(baseInt)
	fmt.Println(float64Int)

	baseInt = 15
	var boolInt bool
	if baseInt != 0 {
		boolInt = true
	} else {
		boolInt = false
	}
	fmt.Println(boolInt)

	baseInt = 15
	byteInt := byte(baseInt)
	fmt.Println(byteInt)

	// Float
	baseFloat64 := 3.1423423

	stringFloat64 := fmt.Sprintf("%f", baseFloat64)
	fmt.Println(stringFloat64)

	baseFloat64 = 3.1423423
	intFloat64 := int(baseFloat64)
	fmt.Println(intFloat64)

	baseFloat64 = 3.1423423
	var boolFloat64 bool
	if baseFloat64 != 0 {
		boolFloat64 = true
	} else {
		boolFloat64 = false
	}
	fmt.Println(boolFloat64)

	baseFloat64 = 3.1423423
	byteFloat64 := byte(baseFloat64)
	fmt.Println(byteFloat64)

	// Bool
	baseBool := true

	stringBool := fmt.Sprintf("%t", baseBool)
	fmt.Println(stringBool)

	baseBool = true
	var intBool int
	if baseBool {
		intBool = 1
	} else {
		intBool = 0
	}
	fmt.Println(intBool)

	baseBool = true
	var float64Bool float64
	if baseBool {
		float64Bool = 1.0
	} else {
		float64Bool = 0.0
	}
	fmt.Println(float64Bool)

	baseBool = true
	var byteBool byte
	if baseBool {
		byteBool = 1
	} else {
		byteBool = 0
	}
	fmt.Println(byteBool)

	// Byte
	baseByte := byte(65)

	stringByte := fmt.Sprintf("%c", baseByte)
	fmt.Println(stringByte)

	baseByte = byte(65)
	intByte := int(baseByte)
	fmt.Println(intByte)

	baseByte = byte(65)
	float64Byte := float64(baseByte)
	fmt.Println(float64Byte)

	baseByte = byte(65)
	var boolByte bool
	if baseByte != 0 {
		boolByte = true
	} else {
		boolByte = false
	}
	fmt.Println(boolByte)
}
