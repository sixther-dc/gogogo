package test

import "fmt"

type height int //身高
type weight int //体重

//给类型增加方法
func (h height) Say(scale int) string {
	return fmt.Sprintf("Height: %d", int(h)*scale)
}
func (w weight) Say() string {
	return fmt.Sprintf("Weight: %d", w)
}

//RunType go_type的主函数
func RunType() {
	fmt.Printf("this is test_type.go\n")

	var x height = 2
	var y weight = 50

	fmt.Printf("%s\n", x.Say(4))
	fmt.Printf("%s\n", y.Say())

}
