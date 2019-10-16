package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//RunURL go_url的主函数
func RunURL() {
	fmt.Println("this is test_url.go")

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch error: %v \n", err)
			os.Exit(1)
		}
		body, err := ioutil.ReadAll(resp.Body)
		//resp.Body是一个可读的服务器响应流
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reding %s error: %v \n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", body)
	}
}
