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

	fmt.Println("*******************")
	s_3 := []int{1, 2, 3, 4, 5}
	fmt.Printf("%p\n", &s_3)
	s_3 = append(s_3[:1], s_3[2:]...)
	fmt.Println(s_3)
	fmt.Printf("s_3的地址%p\n", &s_3)
	// fmt.Println(s_3[2:100]) not ok
	modify(s_3)
	fmt.Println(s_3, cap(s_3))
	fmt.Println("*******************")
	delete(s_3)
	fmt.Printf("s_3的地址%p\n", &s_3)
	fmt.Println("s_3的内容", &s_3)
	fmt.Printf("s_3[0]的地址%p\n", &s_3[0])

}

func modify(nums []int) {
	// nums[0] += 100 // 扩容前修改
	nums = append(nums, 99)
	nums = append(nums, 101)
	nums[0] += 100 // 扩容后修改
	fmt.Println("nums: ", nums)
}

func delete(nums []int) {
	nums = append(nums[:2], nums[3:]...)
	nums[0] = 1001
	fmt.Printf("重组后的切片地址%p\n", &nums)
	fmt.Println("重组后的切片内容", nums)
}

/*
Conclusion:
	1. 无法在不指定数组长度的情况下创建一个数组, 要么[len]type, 要么[]type {value_1, value_2, ... , value_n}
	2. 数组和切片的表观不同在于%#v输出时, 前者带有长度, 而后者没有, 在函数中申明, 确实也是这样区分数组与切片
	3. 但是它们都无法直接扩充长度, 这个过程会复制, 扩充前的的切片如果没有被回收, 是不变的
	4. 切片是一个结构体, 包含数组指针(切片开始位置?), 长度, 和最大长度信息
	5. 切片的容量不等于其相关数组的最大长度, 并且, 切片无法对超出其长度的索引做处理, 会报错, 那知道容量的意义在哪?
	6. 当cap==len时 ,调用append, 添加一个元素, 会使cap变为原来的两倍
	7. 无法这样创建分片 var nums []int {1,2,3,4}
	8. 切片扩容后与原内置数组之间的关系--先扩容后拷贝
	9. 切片的其他操作 https://blog.csdn.net/fyxichen/article/details/46622763
	10. 对slice而言, length很可能是一个属性, 在slice建立之初就创建了, len只是调用这个属性
	11. slice在函数中的传递其实是拷贝, 不过slice中保存了arrya的地址, 所以
	12. 如果将slice中间的元素做一个删除操作, 那么还会是原来的数组么?目前的测试结果是会
	13. 切片中每个元素的地址都是一样的, 但是在上下文中可以通过同样的地址复原出对应的值
	14. (不对)切片的赋值是复制, 这个与python不一样
	15. append(slice, ele_1, ele_2, ...)或者append(slice, slice_1...)

	17. slice向前扩充后, 如果长度没有超过当前的数组容量, 地址会向后移动, 每个整数移动16字节
	18. 拷贝, copy后, 第一个参数(拷贝对象)的长度不会改变
*/
