package method

import (
	"fmt"
	"math"
)

//Run method的主方法
func Run() {
	fmt.Println("this is Run function")
	a := Point{1, 1}
	b := Point{2, 2}
	c := Point{3, 3}
	fmt.Println(a.Distance(b))

	p := Path{a, b, c}
	fmt.Println(p.Distance())

	a.ScaleBy(10)
	fmt.Println(a)
}

//Point 坐标点
type Point struct {
	X, Y float64
}

//Distance  返回两个坐标点之间的距离
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//Path  一条由一些列点组成的线
type Path []Point

//Distance 返回一条线的距离之和
func (p Path) Distance() float64 {
	sum := 0.0
	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}

//ScaleBy 基于指针的方法，将点缩放n倍数
func (p *Point) ScaleBy(n float64) {
	//p为一个指针， *p为指针指向的实际的变量
	fmt.Println(*p)
	(*p).X *= n
	(*p).Y *= n
}
