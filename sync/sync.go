package syncMy

import (
	"fmt"
	"sync"
	"time"
)

func Mutex() {
	mu := sync.Mutex{}
	mu.Lock()
	mu.Unlock()
}

func RWMutex() {
	resource := 0
	//mu := sync.Mutex{}
	mu := sync.RWMutex{}

	go func() {
		for {
			//mu.Lock()
			time.Sleep(1 * time.Second)
			fmt.Println(fmt.Sprintf("reader one: %d", resource))
			//mu.Unlock()
		}
	}()

	go func() {
		for {
			//mu.Lock()
			time.Sleep(3 * time.Second)
			fmt.Println(fmt.Sprintf("reader two: %d", resource))
			//mu.Unlock()
		}
	}()

	go func() {
		for {
			mu.Lock()
			resource++
			mu.Unlock()
			time.Sleep(5 * time.Second)
		}
	}()

	timeS := 0
	for {
		time.Sleep(1 * time.Second)
		timeS++
		fmt.Println(timeS)
	}
}

func WithSlice() {
	type Obj struct {
		Value string
	}
	var a string
	var b Obj

	sliceData := make([]Obj, 0)
	sliceData = append(sliceData, Obj{Value: "1"}, Obj{Value: "2"}, Obj{Value: "3"})

	go func() {
		for {
			b = sliceData[0]
			if b.Value == "askdlkajsdkhasiudhiadh" {
			}
			sliceData = append(sliceData, Obj{Value: "5"})
			//sliceData[0] = Obj{Value: "5"}
		}
	}()

	go func() {
		for {
			b = sliceData[0]
			if b.Value == "askdlkajsdkhasiudhiadh" {
			}
			sliceData = append(sliceData, Obj{Value: "5"})
			//sliceData[0] = Obj{Value: "6"}
		}
	}()
	go func() {
		for {
			for key, value := range sliceData {
				fmt.Println(len(sliceData))
				a = fmt.Sprintf("%d-%s", key, value.Value)
				if a == "askdlkajsdkhasiudhiadh" {
				}
			}
		}
	}()
	for {

	}
}

func WithMap() {
	type Obj struct {
		Value string
	}
	var a string

	mapData := make(map[string]Obj)
	mapData["one"] = Obj{Value: "1"}
	mapData["two"] = Obj{Value: "2"}
	mapData["three"] = Obj{Value: "3"}

	go func() {
		for {
			mapData["four"] = Obj{Value: "4"}
		}
	}()
	go func() {
		for {
			for key, value := range mapData {
				a = fmt.Sprintf("%s-%s", key, value.Value)
				if a == "askdlkajsdkhasiudhiadh" {
				}
			}
		}
	}()
	for {

	}
}
