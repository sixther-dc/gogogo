package function

import (
	"fmt"
	"sort"
)

//Square  可以保持状态计算平方的函数
func Square() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

//AnonFunc 匿名函数的测试函数
func AnonFunc() {
	// a := Square()
	// fmt.Println(a())
	// fmt.Println(a())

	//定义order存储最终的结果
	var order []string
	//定义是否有子节点的flag字典
	seen := make(map[string]bool)

	var visitAll func(items []string)
	//深度优先算法, 深度优先算法最初是为了解决迷宫问题
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(prereqs[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range prereqs {
		keys = append(keys, key)
	}
	//对数组keys进行排序
	sort.Strings(keys)
	fmt.Println(keys)
	visitAll(keys)
	for index, item := range order {
		fmt.Printf("%d\t%s\n", index, item)
	}

}

//如下的每个课程都有自己的前置课程，求一个课程列表，使其可以形成一系列完整的课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}
