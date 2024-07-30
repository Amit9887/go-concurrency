package main

import (
	"fmt"
)

func timesThree(arr []int, ch chan int) {
	for _, elem := range arr {
		ch <- elem * 3
	}
	close(ch)
}

func main() {
	fmt.Println("We are executing a goroutine")
	arr := []int{2, 3, 4, 5, 6}
	ch := make(chan int, len(arr))
	go timesThree(arr, ch)
	for i := range ch {
		fmt.Printf("Result: %v \n", i)
	}
}
