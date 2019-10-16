package main

import (
	"fmt"
	"miffy/method"
	"miffy/test"
)

func main() {
	fmt.Printf("I am main fun of main package\n")
	// test.RunFlag()
	fmt.Printf("%s\n", test.TestVar)
	// fmt.Printf("%s\n", test.TestGifVar)
	fmt.Println("================")
	// 基础类型测试函数
	// datatype.Run()
	// 复杂类型测试函数
	// datatype.RunArrayStruct()
	// 参数的数组传值测试
	// datatype.TestArray()

	// 结构体测试函数
	//datatype.RunStruct()

	//递归，解析html
	// function.ParseMyBlog()
	// function.ParseStructOfMyBlog()
	// function.BeautyStructOfMyBlog()

	//测试匿名函数
	// function.AnonFunc()

	//方法部分的测试
	method.Run()
}
