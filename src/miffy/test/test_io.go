package test

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	//input.scan在读取到新行的时候会返回true，没有会返回false
	for input.Scan() {
		counts[input.Text()]++
	}
}

//RunIo go_io的主函数
func RunIo() {
	fmt.Println("this is test1.go")
	//声明一个key是sting, value为int的map数据结构
	files := os.Args[1:]
	counts := make(map[string]int)
	fmt.Println(files)
	if len(files) < 1 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				//将错误输出到标准错误输出
				fmt.Fprintf(os.Stderr, "test1.go: %v \n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	//对map进行便利的range函数不保序, 每次执行的顺序是随机的
	for line, n := range counts {
		if n > 0 {
			//格式化输出函数
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
