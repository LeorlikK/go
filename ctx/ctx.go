package ctx

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Main() {
	//ctx := context.Background()
	////toDo := context.TODO()
	//
	//withValue := context.WithValue(ctx, "name", "Kirill")
	//fmt.Println(withValue.Value("name"))
	//
	//withCancel, cancel := context.WithCancel(ctx)
	//fmt.Println(withCancel.Err())
	//cancel()
	//fmt.Println(withCancel.Err())
	//
	//withDeadline, cancel := context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	//defer cancel()
	//fmt.Println(withDeadline.Deadline())
	//fmt.Println(withDeadline.Err())
	//fmt.Println(<-withDeadline.Done())
	//
	//withTimeout, cancel := context.WithTimeout(ctx, 3*time.Second)
	//defer cancel()
	//fmt.Println(<-withTimeout.Done())

	timeStart := time.Now()
	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}
	requestSlice := []string{}
	withTimeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			timeAfter := time.After(5 * time.Second)

			select {
			case <-withTimeout.Done():
				return
			case <-timeAfter:
				mu.Lock()
				requestSlice = append(requestSlice, fmt.Sprintf("request_%d", i))
				mu.Unlock()
			}

		}(i, &wg)
	}

	wg.Wait()
	fmt.Println(time.Since(timeStart))
	fmt.Println(requestSlice)
}
