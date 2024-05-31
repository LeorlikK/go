package assert_type

import (
	"fmt"
	"log"
)

func Main() {
	var a interface{} = "test"
	b, ok := a.(string)
	if !ok {
		log.Fatal(ok)
	}
	fmt.Println(b)

	type TestStruct struct {
		Name string
		Age  int
	}
	var c interface{} = &TestStruct{Name: "One", Age: 32}
	d, ok := c.(*TestStruct)
	if !ok {
		log.Fatal(ok)
	}
	fmt.Println(d)

	var arr = []interface{}{
		"Hello, World!",
		42,
		3.14,
		true,
		TestStruct{Name: "Alice", Age: 30},
	}
	for _, v := range arr {
		switch v := v.(type) {
		case string:
			fmt.Printf("Строка: %s\n", v)
		case int:
			fmt.Printf("Целое число: %d\n", v)
		case float64:
			fmt.Printf("Число с плавающей точкой: %f\n", v)
		case bool:
			fmt.Printf("Булево значение: %t\n", v)
		case TestStruct:
			fmt.Printf("Структура Person: {Name: %s, Age: %d}\n", v.Name, v.Age)
		default:
			log.Printf("Неизвестный тип: %T\n", v)
		}
	}
}
