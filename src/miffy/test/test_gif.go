package test

import "fmt"

var testGifVar = "from gif"

//TestGifVar 测试一下
var TestGifVar = "from gif"

//RunGif 可以被外部调用
func RunGif() {
	fmt.Println("this is test_gif.go")
}

//被导入或者自己执行的时候一定会执行init函数
func init() {
	fmt.Printf("I am init function from miffy/test\n")
}

func test() {
	fmt.Printf("I am test function\n")
}
