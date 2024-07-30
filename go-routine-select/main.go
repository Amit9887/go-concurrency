package main

import (
	"fmt"
)

func timesThree(arr []int, ch chan int) {
	for _, elem := range arr {
		ch <- elem * 3
	}
}

func minusThree(arr []int, ch chan int) {
	for _, elem := range arr {
		ch <- elem - 3
	}
}

func main() {
	fmt.Println("We are executing a goroutine")
	arr := []int{2, 3, 4, 5, 6}
	ch := make(chan int, len(arr))
	minusCh := make(chan int, len(arr))

	go timesThree(arr, ch)
	go minusThree(arr, minusCh)

	for i := 0; i < len(arr)*2; i++ {
		select {
		case msg1 := <-ch:
			fmt.Printf("Result timesThree: %v \n", msg1)
		case msg2 := <-minusCh:
			fmt.Printf("Result minusThree: %v \n", msg2)
		default:
			fmt.Println("Non blocking way of listening to multiple channels")
		}
	}
}
