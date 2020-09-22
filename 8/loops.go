package main

import (
	"fmt"
	"time"
)

func main() {
	// numbers := make(chan int)

	// deadlock — когда горутина пытается прочитать значение из канала, но в канале их нету и записи в него не происходит.
	// go generateNumbers(1000, numbers)

	// for {
	// 	number, ok := <-numbers
	// 	fmt.Println(number, ok)
	// 	if !ok {
	// 		break
	// 	}
	// }

	number := make(chan int)

	go func() {
		number <- 42
	}()

	time.Sleep(100 * time.Millisecond)

	select {
	case n := <-number:
		fmt.Println(n)
	default:
		fmt.Println("nothing")
	}

	// for n := range numbers {
	// 	fmt.Println(n)
	// }

}

// func generateNumbers(n int, res chan int) {
// 	for i := 0; i <= n; i++ {
// 		res <- i * 2
// 	}
// 	close(res)
// }
