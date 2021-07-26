package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("原数组: ", s)
	fmt.Println("作为返回值反转: ", IntReverse(s, true))
	fmt.Println("原数组: ", s)
	IntReverse(s, false)
	fmt.Println("就地反转后的原数组: ", s)

}

func IntReverse(s []int, isCopy bool) []int {
	slice := make([]int, len(s))
	fmt.Println("**********")
	fmt.Printf("初始slice地址: %p\n", slice)
	if isCopy {
		copy(slice, s)
		fmt.Printf("复制后slice地址: %p\n", slice)

	} else {
		slice = s
	}
	fmt.Println(slice)

	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[j], slice[i] = slice[i], slice[j]
		fmt.Println(slice)
	}
	if isCopy {
		fmt.Println(slice)
		return slice
	} else {
		fmt.Println(slice)
		return []int{}
	}
	// fmt.Println("yes")

}
