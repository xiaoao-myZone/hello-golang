package main

import (
	"fmt"
	"sort"
)

func main() {
	// 数组
	fmt.Println("********数组***********")
	var arr [8]int //声明长度后, 每个元素自动初始化为0
	fmt.Println(arr)
	fmt.Printf("arr addr is %p\n", &arr)
	initArr1(arr)
	initArr2(&arr)

	var drr = [5]int{1, 2, 3, 4} // 注意这种初始化的不同之处, 从左至右初始化, 不够的保持默认初始值
	fmt.Println(drr)

	// a := 5
	// var prr = [a]int{1, 2, 3, 4} // 注意这种初始化的不同之处, 从左至右初始化, 不够的保持默认初始值
	// fmt.Println(prr)
	//与C一样， 数组初始化长度必须是整形

	fmt.Println("*******************")
	var srr = [...]string{"yes", "no", "hi"} //注意字符串数组的特殊声明方式
	fmt.Println(srr)
	fmt.Printf("%#v\n", srr)

	fmt.Println("*******************")
	var dss = [5]string{2: "sss", 4: "ppp"}
	for i := range dss {
		fmt.Println(dss[i])
	}

	// 切片
	fmt.Println("********切片***********")
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
	fmt.Println("\n********切片的地址问题***********")
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
	fmt.Println("\n********操作切片的接口***********")
	fmt.Printf("ss的地址: %p\n", &ss)
	ss = append(ss, 100)
	ss = append(ss, 101, 102, 103)
	addition := []int{7, 8}
	ss = append(ss, addition...)
	fmt.Printf("append操作后ss的地址: %p\n", &ss)
	fmt.Println(ss)

	// 探究切片的赋值
	fmt.Println("\n********探究切片的赋值***********")
	two_arr := [][]int{{1, 2}, {3, 4}, {5, 6}}
	arr_map := map[int][]int{}
	for _, pairs := range two_arr {
		for _, num := range pairs {
			arr_map[num] = pairs
		}

	}
	fmt.Printf("原数组地址: %p\n", two_arr)
	fmt.Printf("第一个原数组元素地址: %p\n", two_arr[0])
	for k, v := range arr_map {
		fmt.Printf("key: %v  val: %v  addr: %p\n", k, v, v)
	}

	// 切片扩容
	fmt.Println("\n********切片右边扩容***********")
	pp := []int{1}
	ll := pp
	fmt.Printf("原切片pp地址: %p  容量: %v  长度: %v \n", pp, cap(pp), len(pp))
	fmt.Printf("原切片ll地址: %p  容量: %v  长度: %v \n", ll, cap(ll), len(ll))
	for i := 10; i > 0; i-- {
		pp = append(pp, i)
		fmt.Printf("现切片地址: %p  容量: %v  长度: %v \n", pp, cap(pp), len(pp))
	}
	fmt.Printf("原切片ll地址: %p  容量: %v  长度: %v \n", ll, cap(ll), len(ll))
	yy := pp
	fmt.Printf("原切片yy地址: %p  容量: %v  长度: %v \n", yy, cap(yy), len(yy))
	fmt.Println("pp数组增一")
	pp = append(pp, 11)
	fmt.Printf("现pp切片地址: %p  容量: %v  长度: %v \n", pp, cap(pp), len(pp))
	fmt.Printf("现切片yy地址: %p  容量: %v  长度: %v \n", yy, cap(yy), len(yy))
	fmt.Printf("yy: %v\n", yy)
}

func initArr1(arr [8]int) {
	for i := range arr {
		arr[i] = i
	}
	fmt.Printf("传值给函数后地址为:  %p\n", &arr)
}

func initArr2(arr_p *[8]int) {
	fmt.Printf("传址给函数后地址为 %p\n", arr_p)
	// arr := *arr_p 这样是创建(复制)了一个的数组
	for i := range *arr_p {
		fmt.Printf("i: %d\t", i)
		(*arr_p)[i] = i //有没有别的办法引用
	}
	// arr_p -= 1 //不可以像C那样加减指针
	//fmt.Printf("initArr2: arr addr is %p\n", &arr)
}

/*
Conclusion:
	1. 数组更像python的容器类型, 因为它可以存放不同类型的数据(但是一个数组只有一种数据)
	2. 数组初始化后, 长度貌似无法改变


*/
