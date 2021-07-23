package main

import (
	"fmt"
)

func main() {
	var s, i int
	list := []int{1, 2, 3, 4}
	s, i = 5, 0
	s += i
	for j := 0; j < s; j++ {
		fmt.Println("this is circle -->", j)
	}

	for {
		i++
		if i > 20 {
			fmt.Println("loop end")
			break
		}
	}

	for index, val := range list {
		fmt.Println(index, val)
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
	4. for啥也不加相当于while True
	5. for与range结合， 可以很简明地遍历切片与字典
*/
