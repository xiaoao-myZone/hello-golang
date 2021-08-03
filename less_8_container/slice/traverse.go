// go 的range也类似于生成一个迭代器，哪怕切片中有变量， 执行后范围不会有影响

package main

import "fmt"

func main() {
	nums, right := []int{1, 2, 3, 4, 5}, 1
	for right < len(nums) {
		for _, num := range nums[:right] {
			fmt.Println(num)
			right += 1
		}
		fmt.Println("**************")
		// right += 1
	}
}
