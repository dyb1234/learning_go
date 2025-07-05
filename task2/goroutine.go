package main

import (
	"fmt"
	"sync"
	"time"
)

func printNum() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 10; i++ {
			if i%2 == 1 {
				fmt.Println(i)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
	}()
	wg.Wait()
}

func scheduler(tasks []func()) {
	var wg sync.WaitGroup
	for i, elem := range tasks {
		wg.Add(1)
		go func(id int, function func()) {
			defer wg.Done()
			start := time.Now()
			function()
			duration := time.Since(start)
			fmt.Printf("task %d last %v duration\n", id+1, duration)
		}(i, elem)
	}
	wg.Wait()
}
