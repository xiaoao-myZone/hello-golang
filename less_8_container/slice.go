package main

import "fmt"

func main() {
	s_1 := []int{1, 2, 3, 4, 5, 6}
	var a_1 [6]int
	s_2 := a_1[:]
	a_1[2] = 100
	fmt.Printf("%#v\n", s_1)
	fmt.Printf("%#v\n", a_1)
	fmt.Printf("%#v\n", s_2)
	fmt.Println("s_1", s_1)
	fmt.Println("s_2", s_2)
	s_2 = a_1[2:5] //切片变量对象创建后可以直接赋值
	fmt.Println("s_2", s_2)
	fmt.Println("len of s_2:", len(s_2)) // len is 3
	fmt.Println("cap of s_2:", cap(s_2)) // cap is 4
	for i := 0; i < len(s_2); i++ {
		fmt.Println(s_2[i])
	}
	// fmt.Println(s_2[3]) //会报错--超出范围
	s_2 = append(s_2, 0, 255)
	fmt.Println(s_2)
	fmt.Println("len of s_2:", len(s_2)) // len is 5
	fmt.Println("cap of s_2:", cap(s_2)) // cap is 8
	s_2 = append(s_2, 77, 88, 99, 10)
	fmt.Println("len of s_2:", len(s_2)) // len is 5
	fmt.Println("cap of s_2:", cap(s_2)) // cap is 8
	fmt.Println("s_1[:0]:", s_1[:1])
	fmt.Println(s_2)
	s_2 = append(s_2, s_1[:0]...)
	fmt.Println(s_2)
}

/*
Conclusion:
	1. 无法在不指定数组长度的情况下创建一个数组, 要么[len]type, 要么[]type {value_1, value_2, ... , value_n}
	2. 数组和切片的表观不同在于%#v输出时, 前者带有长度, 而后者没有, 在函数中申明, 确实也是这样区分数组与切片
	3. 但是它们都无法直接扩充长度
	4. 切片是一个结构体, 包含数组指针(切片开始位置?), 长度, 和最大长度信息
	5. 切片的容量不等于其相关数组的最大长度, 并且, 切片无法对超出其长度的索引做处理, 会报错, 那知道容量的意义在哪?
	6. 当cap==len时 ,调用append, 会使cap变为原来的两倍
	7. 无法这样创建分片 var nums []int {1,2,3,4}
*/
