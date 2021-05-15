package main

import (
	"fmt"
)

func main() {
	var s, i int
	s, i = 5, 0
	s += i
	for j := 0; j < s; j++ {
		fmt.Println("this is circle -->", j)
	}
	// else:
	// 	fmt.Println("end else")
	// fmt.Println(j)
}

/*
Conclusion:
	1. 整体上和C差不多, 指示少了一对括号
	2. go 继承C的精髓, ++和=+
	3. 不能像python一样用for-else, 略微遗憾
*/
