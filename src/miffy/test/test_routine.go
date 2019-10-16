package test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//RunRoutine go_routine的主函数
func runRoutine() {
	fmt.Printf("this is test_routine.go\n")
	start := time.Now()
	//channel之间传递的内容是string
	ch := make(chan string)
	//启动并发线程
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	//获子线程的输出
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2f total\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		//跟channel通信, 此处不能sed  string之外的东西
		ch <- "error"
		ch <- fmt.Sprint(err)
		return
	}
	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while readind %s: %v\n", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f  %7d  %s", secs, nBytes, url)
}
