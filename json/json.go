package json

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
)

type User struct {
	Name   string `json:"name"`
	Family string `json:"family"`
	age    int    `json:"age"`
}

func Main() {
	//data := "test_string"
	//data := []string{"A", "B", "C"}
	//data := map[string]int{"A": 1, "B": 2, "C": 3}
	//data := map[string]interface{}{"A": "1", "B": 2, "C": true, "D": []string{"a", "b", "c"}}
	data := User{Name: "Kirill", Family: "Orlov", age: 97}

	bate, _ := json.Marshal(data)

	//fmt.Println(string(bate))

	//var unData string
	//var unData []string
	//var unData map[string]int
	//var unData map[string]interface{}
	var unData User

	if err := json.Unmarshal(bate, &unData); err != nil {
		panic("error: Unmarshal")
	}
	fmt.Println(unData)

	GJson()
}

func GJson() {
	jsonString := `{"name": "Kirill", "family": "Orlov", "info": {"age": 97}}`
	info := gjson.Get(jsonString, "info.age")
	fmt.Println(info)
}
