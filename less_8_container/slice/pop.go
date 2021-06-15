//探究将slice中的一个元素弹出会对原有的slice有何影响

package main

import "fmt"

func main() {
	nums := []int{1, 2}
	fmt.Println(nums)
	new := append(nums[:0], nums[1:]...)
	fmt.Println(nums, new) //nums由[1, 2]变为[2, 2]
}
