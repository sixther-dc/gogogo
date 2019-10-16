package datatype

import (
	"fmt"
	"math"
	"math/cmplx"
	"unicode/utf8"
)

//Run datatype包的主入口
func Run() {
	//int8 取值-128 到 127
	var a int8
	a = 7
	//如果在for循环中使用无符号的类型时, i>=0 将永远成立
	for i := int8(a); i >= 0; i-- {
		fmt.Println(i)
	}
	fmt.Printf("%f\n", math.MaxFloat32)
	fmt.Println(1i * 1i) //-1
	//求一个负数的开方
	fmt.Println(cmplx.Sqrt(-5))
	// fmt.Printf("%f\n", math.MaxFloat64)
	//s中的一个字符包含两个byte
	var s string = "hello,段超"
	fmt.Println(len(s)) //12
	//根据utf8的结构统计字符
	fmt.Println(utf8.RuneCountInString(s)) //8
	//默认以byte为单位获取字符
	fmt.Printf("%c\n", s[0])
	for i := 0; i < len(s); {
		//使用DecodeRuneInString解析utf8中的字符
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	// go中的range默认是用utf8进行解析
	for i, r := range s {
		fmt.Printf("%d\t%c\n", i, r)
	}

	fmt.Println(basename("a/b/a.d.go"))
	fmt.Println(comma("1000000000"))
}

//HasPrefix 前缀测试
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

//HasSuffix 后缀测试
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[(len(s)-len(suffix)):] == suffix
}

/*
/a/b/c.go  --> c
/a/b/c ---> c
/a/b/a.d.go ---> a.d
*/
func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

/*
每三位增加逗号
100000000    ->  1,000,000,000
使用迭代解决
*/
func comma(s string) string {
	n := len(s)
	if n < 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
