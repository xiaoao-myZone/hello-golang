package main

import "fmt"

func main() {
	// map_a := map[string]int{"a": 5}
	var map_a = map[string]int{}
	//map_a = map[string]int{"hh": 55}
	map_a["b"] = 4
	fmt.Println(map_a)
	map_b := make(map[string]int) //即使make会自动给map初始化为一个空列表
	fmt.Println(map_b)
	m, n := fmt.Println(map_a["a"])
	s, ok := map_a["a"]
	fmt.Println(s, ok)
	fmt.Println(m, n)

	// -----------------
	iter_m := map[int]map[int]int{}
	//iter_m[1][2] = 0
	fmt.Println(iter_m[1])
}

/*
Conclusion:
	1. map做参数传递所消耗的内存很少, 4(32bit)或8(64bit)
	2. map的检索的速度是数组或切片的百分之一
	3. map的键与值的特点与python中相同, 键是可hash的对象, 值可以是各种对象
	4. 申明map后, map是nil, 不能直接进行赋值操作, 需要初始化, 这很不方便(可以用{}初始化)
	5. map的增删改查 TODO https://blog.csdn.net/lengyue1084/article/details/105955629
*/
