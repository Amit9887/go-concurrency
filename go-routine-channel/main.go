package main

import (
	"fmt"
	"time"
)

func timesThree(number int, ch chan int) {
	result := number * 3
	ch <- result
}

func main() {
	fmt.Println("We are executing a goroutine")
	ch := make(chan int)
	go timesThree(3, ch)
	result := <-ch
	fmt.Println(result)
	fmt.Println("Done!")
	time.Sleep(time.Second)
}
