package goroutines

import (
	"fmt"
	"time"
)

func Main() {
	ch := make(chan int)
	fmt.Println("routine:start")
	go say("Hello World!", ch)
	fmt.Println("routine:finish")
	fmt.Println(<-ch)
}

func say(text string, ch chan int) {
	time.Sleep(5 * time.Second)
	fmt.Println(text)
	ch <- 1
}
