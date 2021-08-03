package main

import "fmt"

func main() {
	// map_a := map[string]int{"a": 5}
	var map_a = map[string]int{}
	//map_a = map[string]int{"hh": 55}
	map_a["b"] = 4
	map_a["a"] = 5
	fmt.Println(map_a)
	map_b := make(map[string]int) //即使make会自动给map初始化为一个空列表
	fmt.Println(map_b)
	delete(map_a, "a")
	m, n := fmt.Println(map_a["a"])
	s, ok := map_a["a"]
	fmt.Println(s, ok)
	fmt.Println(m, n)

	// -----------------
	iter_m := map[int]map[int]int{}
	//iter_m[1][2] = 0
	fmt.Printf("%#v\n", iter_m[1])
	iter_s := map[int][3]int{}
	fmt.Printf("%#v\n", iter_s[1])
	iter_p := map[int]sss{}
	fmt.Printf("%#v\n", iter_p[1])

	fmt.Println("********作为参数********")
	k := map[int]int{1: 2, 4: 1}
	fmt.Println("修改前", k)
	change_map(k)
	fmt.Println("修改后", k)

}

type sss struct {
	mm int
	kk []int
}

func change_map(m map[int]int) {
	m[4] = 100
}

/*
Conclusion:
	1. map做参数传递所消耗的内存很少, 4(32bit)或8(64bit)
	2. map的检索的速度是数组或切片的百分之一
	3. map的键与值的特点与python中相同, 键是可hash的对象, 值可以是各种对象
	4. 申明map后, map是nil, 不能直接进行赋值操作, 需要初始化, 这很不方便(可以用{}初始化)
	5. map的增删改查 TODO https://blog.csdn.net/lengyue1084/article/details/105955629
	6. delete(map, key)函数可以删除键值对
	7. 如果map没有一个键, 那么搜索的结果是什么?  根据情况, 如果是切片和字典, 返回的是nil, 如果是数组, 将会返回一个初始化的数组(全0)
	   如果是结构体, 将会返回一个全部是初始值的结构体实例
	8. map作为值传递是传值还是传址? 传址
*/
