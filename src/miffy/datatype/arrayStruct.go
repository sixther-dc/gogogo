package datatype

import "fmt"

//RunArrayStruct  复杂结构函数主入口
func RunArrayStruct() {
	a := [3]int{1, 2, 3}
	var b *[3]int
	b = &a
	fmt.Println(a[0])
	fmt.Println(zeroArray(b))

	c := make([]int, 3, 4)
	fmt.Println(c)
	c = append(c, 4)
	fmt.Println(c)
	fmt.Println(remove(c, 3))

	ages := make(map[string]int)
	ages["duanchao"] = 29
	ages["simon"] = 25
	fmt.Printf("%v\n", ages)

}

//指针传值，数组只能处理特定长度的数组
func zeroArray(ptr *[3]int) [3]int {
	*ptr = [3]int{}
	return *ptr
}

//删除数组某一项
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
