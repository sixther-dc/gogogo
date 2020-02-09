package routine

import (
	"fmt"
	"os"
	"time"
)

//select每次一定要选择一个case来执行，只要有ch有数据过来，第一个过来的channal会得到处理
func testSelect() (interface{}, string) {
	ch1 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "errof"
	}()
	//只要有channel接收数据就会return, 类似RXjs
	select {
	case err := <-ch1:
		close(ch1)
		return nil, err
	case <-time.After(3 * time.Second):
	}
	fmt.Println("duanchao")
	return nil, "test"
}

func testTick() {
	tick := time.Tick(1 * time.Second)
	abort := make(chan struct{})
	inputBytes := make([]byte, 3)
	go func() {
		//遇到回车返回, 获取标准输入
		n, _ := os.Stdin.Read(inputBytes)
		fmt.Printf("..............%v\n", string(inputBytes))
		fmt.Printf("..............%d\n", n)
		abort <- struct{}{}
	}()
	for i := 10; i > 0; i-- {
		select {
		case <-tick:
			fmt.Println(i)
		case <-abort:
			fmt.Println("aborted...")
			return
			// default:
			// 	fmt.Println("default")
		}

	}

	fmt.Println("lanuch....")
	return
}

//Run select主测试方法
func Run() {
	// _, err := testSelect()
	// println(err)
	// println("====")
	testTick()
}
