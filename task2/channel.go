package main

import (
	"fmt"
	"sync"
)

func channel1() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func(c chan int) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			c <- i + 1
		}
	}(ch)
	go func(c chan int) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
	}(ch)
	wg.Wait()
}
func channel2() {
	ch := make(chan int, 100)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(c chan int) {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			c <- i + 1
		}
		close(c)
	}(ch)
	wg.Wait()
	wg.Add(1)
	go func(c chan int) {
		defer wg.Done()
		for elem := range c {
			fmt.Print(elem)
			fmt.Print(" ")
		}
		fmt.Println()
	}(ch)
	wg.Wait()
}
