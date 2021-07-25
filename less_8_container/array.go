package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr [8]int

	fmt.Printf("arr addr is %p\n", &arr)
	initArr1(arr)
	initArr2(&arr)
	// arr[8] = 100 长度不可以更改
	for i := range arr {
		fmt.Println(arr[i])
	}

	fmt.Println("*******************")
	var drr = [5]int{1, 2, 3, 4} // 注意这种初始化的不同之处, 从左至右初始化, 不够的保持默认初始值
	for i := range drr {
		fmt.Println(drr[i])
	}

	fmt.Println("*******************")
	var srr = [...]string{"yes", "no", "hi"}
	for i := range srr {
		fmt.Println(srr[i])
	}

	fmt.Println("*******************")
	var dss = [5]string{2: "sss", 4: "ppp"}
	for i := range dss {
		fmt.Println(dss[i])
	}

	slides := arr[:] // 这个切片操作和python相同, 如果不存在将返回空切片, 前提是:前后的索引不可以超出最大长度
	fmt.Println(slides)

	var d_arr [2][2]int
	for i := range d_arr {
		fmt.Println("i-->", i)
		for j := range d_arr[i] {
			fmt.Println("j-->", d_arr[i][j])
		}
	}

	fmt.Println("*******************")
	ss := []int{5, 3, 4, 1}
	sort.Ints(ss)
	fmt.Println(ss)

	// 切片的地址问题
	fmt.Println("********切片的地址问题***********")
	new := []int{1, 2, 3}
	new_p := []*int{}
	for _, i := range new {
		new_p = append(new_p, &i)
		fmt.Printf("%p\n", &i)
		fmt.Println(*(&i)) // 仿佛这个地址受上下文的影响
	}
	for _, i := range new_p {
		fmt.Println(*i)
	}

	// 操作切片的接口
	fmt.Println("********操作切片的接口***********")
	fmt.Printf("ss的地址: %p\n", &ss)
	ss = append(ss, 100)
	ss = append(ss, 101, 102, 103)
	addition := []int{7, 8}
	ss = append(ss, addition...)
	fmt.Printf("append操作后ss的地址: %p\n", &ss)
	fmt.Println(ss)

}

func initArr1(arr [8]int) {
	for i := range arr {
		arr[i] = i
	}
	fmt.Printf("initArr1: arr addr is %p\n", &arr)
}

func initArr2(arr_p *[8]int) {
	fmt.Printf("initArr2: arr addr is %p\n", arr_p)
	// arr := *arr_p 这样是创建(复制)了一个的数组
	for i := range *arr_p {
		fmt.Printf("i: %d\t", i)
		(*arr_p)[i] = i //有没有别的办法引用
	}
	//fmt.Printf("initArr2: arr addr is %p\n", &arr)
}

/*
Conclusion:
	1. 数组更像python的容器类型, 因为它可以存放不同类型的数据(但是一个数组只有一种数据)
	2. 数组初始化后, 长度貌似无法改变
	3. 切片中每个元素的地址都是一样的, 但是在上下文中可以通过同样的地址复原出对应的值
	4. 切片的赋值是复制, 这个与python不一样
	5. append(slice, ele_1, ele_2, ...)或者append(slice, slice_1...)
*/
