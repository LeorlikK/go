package goroutines

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func Main() {
	// ch := make(chan int) - создание
	// <- ch - чтение из канала
	// ch <- - запись/вывод из канала
	// for i:= range ch {
	// fmt.Println(i)
	// }
	// close(ch)

	time1 := time.Now()

	response := AllAsyncRequest()

	time2 := time.Now()
	time3 := time2.Sub(time1)

	fmt.Println(response)
	fmt.Println(time3)
}

//	func say(text string, ch chan int) {
//		time.Sleep(5 * time.Second)
//		fmt.Println(text)
//		ch <- 1
//	}
func jsonParse(response *http.Response) []string {
	defer requestClose(response)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var result []map[string]interface{}
	data := make([]string, 0)

	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}

	for _, value := range result {
		address, err := value["address"].(map[string]interface{})
		if !err {
			return []string{"Error: 1"}
		} else {
			street, err := address["street"].(string)
			if !err {
				return []string{"Error: 2"}
			} else {
				data = append(data, street)
			}
		}
	}

	return data
}

func AllAsyncRequest() [][]string {
	var allData [][]string
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		ch := make(chan int)
		go func(i int, ch chan int) {
			defer wg.Done()

			response := Request()
			data := jsonParse(response)
			allData = append(allData, data)
		}(i, ch)
	}

	wg.Wait()
	return allData
}

func AllSyncRequest() [][]string {
	var allData [][]string

	for i := 0; i < 50; i++ {
		response := Request()
		data := jsonParse(response)
		allData = append(allData, data)
	}

	return allData
}

func Request() *http.Response {
	response, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		requestClose(response)
		panic(err)
	}

	return response
}

func requestClose(response *http.Response) {
	response.Body.Close()
}
