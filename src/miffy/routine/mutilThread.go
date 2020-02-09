package routine

import (
	"fmt"
	"sync"
	"time"
)

//RunBingfa fun
func RunBingfa() {
	var a = []int{1, 2, 3, 4}
	var wg sync.WaitGroup
	var result = make(chan int)
	for _, i := range a {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			time.Sleep(time.Second * time.Duration(x))
			result <- x
		}(i)
	}

	//等待计数器清零
	go func() {
		wg.Wait()
		close(result)
	}()

	var total int
	for count := range result {
		total += count
	}

	fmt.Println(total)
}
