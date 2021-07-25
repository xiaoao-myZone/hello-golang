package main

import "fmt"

const (
	x int = 5 + iota
	y
	z
)

type company struct {
	employee_num int
	set_up_data  string
}

func main() {
	var i = 5
	var j int = 4
	var p *int = nil             // nil竟是一个关键字
	fmt.Printf("%d, %p\n", i, p) // 居然是 0x0
	fmt.Println(i, j, p, x, y, z)
	fmt.Printf("%+v\n", p) // %+v输出 0xc0000140d8 %#v 输出 (*int)(0xc0000140d8)
	//fmt.Println(*p)

	// structure
	fmt.Println("********structure*********")
	c := company{employee_num: 100, set_up_data: "2018"}
	fmt.Println("c.employee_num: ", c.employee_num)
	pc := &c
	fmt.Printf("初始化后的结构体实例地址: %p\n", &c)
	fmt.Println("c.employee_num: ", pc.employee_num)
	point_test_1(c)
	point_test_2(pc)

	// point operation
	// s := []int{1, 2, 3, 4}
	// ps := &s
	// fmt.Printf(ps + 1) //error
}

func point_test_1(c company) {
	fmt.Printf("传值给函数后的结构体实例地址: %p\n", &c)
	fmt.Println("This is test 1")
	fmt.Println("c.employee_num: ", c.employee_num)
	fmt.Println("c.set_up_data: ", c.set_up_data)
}

func point_test_2(c *company) {
	fmt.Printf("传指针给函数后的结构体实例地址: %p\n", c)
	fmt.Println("This is test 2")
	fmt.Println("c.employee_num: ", c.employee_num)
	fmt.Println("c.set_up_data: ", c.set_up_data)
}

/*
Conclusion:
	1. 无论是结构体实例, 还是结构体实例指针, 均可以直接用s.val, C语言中是->可以通过
	2. 即便函数参数变量申明是用结构体指针, 在使用的时候可以直接用丢给它结构体实例
	3. 参数定义为结构体指针, 在函数结构体实例的地址与外界一样, 是传地址
	4. 虽然存在指针, 但是不能像C一样对指针进行加减操作
*/
