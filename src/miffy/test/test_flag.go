package test

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

/**
	go test_flag -sep="#" 123 345
	--> 123#345
**/

//TODO: 这里的sep n为啥需要时数组
var n = flag.Bool("n", false, "omit triling newline")
var sep = flag.String("sep", " ", "separator")

//TestVar 测试变量
var TestVar = "flag"

//RunFlag go_flag的主函数
func RunFlag() {
	fmt.Printf("this is test_flag.go\n")
	flag.Parse()
	*sep = "XXX"
	fmt.Println(strings.Join(os.Args, *sep))
	fmt.Println(strings.Join(flag.Args(), *sep))
	fmt.Println(TestGifVar)
	// test()
	// runGif()
	// RunPoint()
}
