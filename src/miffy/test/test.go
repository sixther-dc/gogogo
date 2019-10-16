package test

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello, world", "rser")
	fmt.Println(os.Args)
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Println(len(os.Args))
	fmt.Println(os.Args[1:len(os.Args)])
	var sep string
	sep = "duanchao"
	fmt.Println(sep)
	// fmt.Println(os.Args[1:])
	for _, arg := range os.Args[1:] {
		//_不能被使用，range返回index，value，称之为空白处理符
		sep += arg
	}
	fmt.Println(sep)
}
