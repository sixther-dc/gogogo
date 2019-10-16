package datatype

import (
	"fmt"
)

func TestArray() {
	//对于一个有元素的数组，函数的参数为引用传递
	a := []int{1,2,3,4,5}
	fmt.Println(a)
	modify(a)
	fmt.Println(a)
	
	//情况数组
	zero(a)
	fmt.Println(a)


	//测试空slice
	var b []int
	modify(b)

}

func modify(a []int) {
	a[0] = 99
	fmt.Println(a)
}

func zero(a []int) {
	a = []int{}
}