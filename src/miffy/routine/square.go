package routine

import (
	"fmt"
)

//RunSquare func 类似于Rxjs的管道
func RunSquare() {
	square := make(chan int)
	result := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			square <- i
			// time.Sleep(1 * time.Second)
		}
		close(square)
	}()

	go func() {
		// for {
		// 	x := <-square
		// 	result <- x * x
		// }
		for x := range square {
			result <- x * x
		}
		close(result)
	}()

	for result := range result {
		fmt.Println(result)
	}
}
