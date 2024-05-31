package test

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var requestCount = 10000000
var chCountRequest = make(chan struct{}, 100)
var data = Data{make(map[string]*ProxyInfo)}

type Data struct {
	ProxyData map[string]*ProxyInfo
}
type ProxyInfo struct {
	timestamp int64
}

func SyncTest() {
	timeStart := time.Now()
	Channel()
	fmt.Println(time.Since(timeStart))

	timeStart = time.Now()
	Mutex()
	fmt.Println(time.Since(timeStart))

	timeStart = time.Now()
	Mutex()
	fmt.Println(time.Since(timeStart))

	timeStart = time.Now()
	Channel()
	fmt.Println(time.Since(timeStart))

	timeStart = time.Now()
	ManyChannel()
	fmt.Println(time.Since(timeStart))

	timeStart = time.Now()
	ManyChannel2()
	fmt.Println(time.Since(timeStart))
}

func Mutex() {
	mu := sync.RWMutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < requestCount; i++ {
		wg.Add(1)
		chCountRequest <- struct{}{}

		go func(i int, data *Data, mu *sync.RWMutex, wg *sync.WaitGroup) {
			defer func() {
				wg.Done()
				<-chCountRequest
			}()

			//time.Sleep(1 * time.Second)

			info := &ProxyInfo{}
			info.timestamp = rand.Int63()
			key := fmt.Sprintf("player_%d", i)

			mu.Lock()
			data.ProxyData[key] = info
			mu.Unlock()

		}(i, &data, &mu, &wg)
	}
	wg.Wait()
	fmt.Println(fmt.Sprintf("Mutex: %d", len(data.ProxyData)))
}
func ManyMutex() {
}

func Channel() {
	done := make(chan struct{}, 100)
	done2 := make(chan struct{})
	var chData = make(chan func(data *Data), 100)

	go func() {
		//for {
		//	select {
		//	case f, ok := <-chData:
		//		if !ok {
		//			return
		//		}
		//		f(&data)
		//	}
		//}
		for f := range chData {
			f(&data)
		}
		done2 <- struct{}{}
	}()

	go func() {
		for i := 0; i < requestCount; i++ {
			chCountRequest <- struct{}{}

			go func(i int, data *Data, done chan struct{}, chData chan func(data *Data)) {
				defer func() {
					done <- struct{}{}
					<-chCountRequest
				}()

				info := &ProxyInfo{}
				info.timestamp = rand.Int63()
				key := fmt.Sprintf("player_%d", i)

				chData <- func(data *Data) {
					data.ProxyData[key] = info
				}
			}(i, &data, done, chData)
		}
	}()

	for i := 0; i < requestCount; i++ {
		<-done
	}

	close(chData)
	<-done2
	fmt.Println(fmt.Sprintf("Channel: %d", len(data.ProxyData)))
}

func ManyChannel() {
	done := make(chan struct{}, 100)
	done2 := make(chan struct{}, 5)
	var chData = make(chan func(data map[string]*ProxyInfo), 100)
	var chSumData = make(chan map[string]*ProxyInfo)

	go func() {
		for i := 0; i < 5; i++ {
			//dataS := <-chSumData
			<-chSumData
			//for key, proxyInfo := range dataS {
			//	data.ProxyData[key] = proxyInfo
			//}
			done2 <- struct{}{}
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			go func() {
				localData := make(map[string]*ProxyInfo)
				for f := range chData {
					f(localData)
				}
				chSumData <- localData
			}()
		}
	}()

	go func() {
		for i := 0; i < requestCount; i++ {
			chCountRequest <- struct{}{}

			go func(i int, done chan struct{}, chData chan func(data map[string]*ProxyInfo)) {
				defer func() {
					done <- struct{}{}
					<-chCountRequest
				}()

				info := &ProxyInfo{}
				info.timestamp = rand.Int63()
				key := fmt.Sprintf("player_%d", i)

				chData <- func(data map[string]*ProxyInfo) {
					data[key] = info
				}
			}(i, done, chData)
		}
	}()

	for i := 0; i < requestCount; i++ {
		<-done
	}
	close(chData)

	for i := 0; i < 5; i++ {
		<-done2
	}

	fmt.Println(fmt.Sprintf("ManyChannel: %d", len(data.ProxyData)))
}

func ManyChannel2() {
	o := StructOne{Size: 2}
	op := &o

	//o.ScaleOne()
	//fmt.Println(o.Size)
	o.ScaleTwo()
	//fmt.Println(o.Size)

	fmt.Println("--------------------")

	op.ScaleOne()
	op.ScaleTwo()
	op.ScaleOne()
}

type StructOne struct {
	Size int
}

func (s StructOne) ScaleOne() {
	s.Size *= s.Size
	fmt.Println(s.Size)
}
func (s *StructOne) ScaleTwo() {
	s.Size *= s.Size
	fmt.Println(s.Size)
}
