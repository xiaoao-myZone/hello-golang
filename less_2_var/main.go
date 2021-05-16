package main

import "fmt"

func main() {
	var num = 0          // 1
	chars := "xiaooa"    // 2
	var ok bool          // 3
	ok = false           // 3
	var f float64 = 5.24 // 4
	var a, b int         // 5
	a, b = 1, 2          // 5
	var c, d = 3, 4      // 6
	a, b = b, c          // same as python
	var (
		m = 5
		n = "havana" // 批量定义
	)
	fmt.Println(a, b, c, d, m, n)
	fmt.Printf("type of num is %T\t type of chars is %T\n", num, chars)
	fmt.Printf("%+v\t%#v\n", ok, ok)
	fmt.Printf("%+v\t%#v\n", f, f)
	fmt.Printf("%+v\t%#v\n", chars, chars) //xiaoao		"xiaoao"
	fmt.Printf("%v\n", chars)              // xiaoao, so %+v is same as %v
	// TODO 类型间如何转化?
}

/*
Conclusion:
	1. %#v 可以输出变量的类型
	2. %v与%+v感觉差不多
*/
