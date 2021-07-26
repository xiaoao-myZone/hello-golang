package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	var s2 []int
	s3 := []int{}
	s4 := make([]int, 7)
	fmt.Printf("s1: %#v\n", s1)
	fmt.Printf("s2: %#v\n", s2)
	fmt.Println("通过var s2 []int创建的slice是不是nil: ", s2 == nil)
	fmt.Printf("s3: %#v\n", s3)
	fmt.Println("通过s3 := []int{}创建的slice是不是nil: ", s3 == nil)
	fmt.Printf("s4: %#v\n", s4)
	fmt.Printf("由make加上长度参数创建的切片的容量是: %+v 长度是: %v\n", cap(s4), len(s4))
	s2 = append(s2, 1)
	fmt.Printf("append后, s2: %#v\n", s2)
}

/*
Conclusion:
	1. var s2 []int只声明不初始化的切片实际就是nil, 打印出来和空切片一样
*/
