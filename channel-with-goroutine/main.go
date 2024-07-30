package main

import (
	"fmt"
	"sync"
)

func timesThree(arr []int, ch chan int, wg *sync.WaitGroup) {
	minusCh := make(chan int, 3)
	for _, elem := range arr {
		value := elem * 3
		if value%2 == 0 {
			wg.Add(1)
			go minusThree(value, minusCh, wg)
			value = <-minusCh
		}
		ch <- value
	}
	close(ch)
}

func minusThree(number int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- number - 3
	fmt.Println("The function continues after returning the result")
}

func main() {
	fmt.Println("We are executing a goroutine")
	arr := []int{2, 3, 4}
	ch := make(chan int, len(arr))
	var wg sync.WaitGroup
	go timesThree(arr, ch, &wg)
	wg.Wait()
	for result := range ch {
		fmt.Printf("Result: %v \n", result)
	}
}
