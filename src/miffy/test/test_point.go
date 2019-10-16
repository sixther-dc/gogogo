package test

import "fmt"

//RunPoint go_point的主函数
func RunPoint() {
	fmt.Printf("this is test_point.go\n")
	var x int = 10
	var y *int
	y = &x
	//y保存了变量x的内存地址
	fmt.Printf("%v\n", y)
	fmt.Printf("%d\n", *y)
	//*y == x   &x == y
	*y = 200
	fmt.Printf("%d\n", x)

	fmt.Printf("%v\n", f())
}

func f() *int {
	var x = 1
	return &x
}
