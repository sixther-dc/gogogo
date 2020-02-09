package routine

import (
	"fmt"
	"time"
)

//RunSpinner 方法
//RunSpinner函数结束后goroutine也会随之结束
func RunSpinner() {
	ch := make(chan int)
	go spinner(100*time.Millisecond, ch)
	go func() {
		for {
			ch <- 101
		}
	}()
	// \r为回到行首
	fmt.Printf("\rFibonacci(%d) = %d\n", 10, fib(40))
	ch <- 100
}

//TODO: 如果通过管道来停止spinner
func spinner(delay time.Duration, ch chan int) {
	for {
		rr := <-ch
		if rr == 100 {
			fmt.Printf("%s\r", "  ")
			break
		}
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
