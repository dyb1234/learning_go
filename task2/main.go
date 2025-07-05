package main

import "fmt"

func main() {
	// pointer
	num := 10
	plus10(&num)
	fmt.Println(num)

	slice1 := []int{1, 2, 3, 4}
	multiply2(&slice1)
	fmt.Println(slice1)

	//goroutine
	printNum()
	funcSlice := []func(){
		func() {
			println(1)
		},
		func() {
			println(2)
		},
		func() {
			println(3)
		},
	}
	scheduler(funcSlice)

	// object oriented programming
	rect := Rectangle{1.5, 2.0}
	fmt.Println(rect.Area(), rect.Perimeter())
	circle := Circle{3}
	fmt.Println(circle.Area(), circle.Perimeter())

	e := Employee{Person{"Bob", 18}, 1234567}
	e.PrintInfo()

	// channel
	channel1()
	channel2()

	//lock
	lock1()
	lock2()
}
