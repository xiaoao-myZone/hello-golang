/*
 */
package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	ret := sum(nums)
	fmt.Println("The result is ", ret)
}

func sum(nums []int) int {
	ret := 0
	for _, i := range nums {
		ret += i
	}
	return ret
}
