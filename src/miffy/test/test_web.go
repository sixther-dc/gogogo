package test

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex

//count在各个携程之间共享，通过sync.Mutex解决竞态条件问题
var count int

//RunWeb go_web的主函数
func RunWeb() {
	fmt.Printf("this is test_web.go\n")

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Fprintf(w, "count = %d\n", count)
}
