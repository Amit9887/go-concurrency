// first run this

// package main

// import (
// 	"fmt"
// 	"time"
// )

// var n = 1

// func timesThree() {
// 	n *= 3
// 	fmt.Println(n)
// }

// func main() {
// 	fmt.Println("We are executing a goroutine")
// 	for i := 0; i < 10; i++ {
// 		go timesThree()
// 	}
// 	time.Sleep(time.Second)
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

var n = 1
var mu sync.Mutex

func timesThree() {
	mu.Lock()
	defer mu.Unlock()
	n *= 3
	fmt.Println(n)
}

func main() {
	fmt.Println("We are executing a goroutine")
	for i := 0; i < 10; i++ {
		go timesThree()
	}
	time.Sleep(time.Second)
}
