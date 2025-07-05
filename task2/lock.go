package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func lock1() {
	counter := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(mup *sync.Mutex) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mup.Lock()
				counter++
				mup.Unlock()
			}
		}(&mu)
	}
	wg.Wait()
	fmt.Println(counter)
}

func lock2() {
	var counter int64 = 0
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter)

}
