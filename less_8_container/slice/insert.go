package main

import "fmt"

func main() {
	// 切片头部插入地址是否会变
	fmt.Println("\n********切片头部插入地址是否会变***********")
	m := []int{1, 2, 3}
	fmt.Printf("原数组地址: %p\n", m)
	n := []int{0}
	fmt.Printf("操作数组地址: %p\n", n)
	m = append(n, m...)
	fmt.Printf("现数组地址: %p, %v\n", m, m)
}

/*
Conclusion:
	1. slice向后扩充后, 会换一个新的地址, 之前的数组会根据引用数量是否为0来回收
*/
