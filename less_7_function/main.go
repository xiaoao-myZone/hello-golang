package main

import (
	"fmt"
	"log"
	"runtime"
)

var num int = 10

func where() {
	_, file_1, file_2, _ := runtime.Caller(1)
	log.Println(file_1, file_2)
}

func main() {
	var k, v, s int
	k, v = getTwoNum(num)
	k, v, s = getTriNum()
	fmt.Println(k, v, s)
	m, n := getMap()
	fmt.Println(m, n)
	p, q := getSlice()
	fmt.Println(p, q)
	fmt.Printf("%#v\n", p)
	if p == nil {
		fmt.Println("It's nil")
	} else {
		fmt.Println("It's not nil")
	}

	fmt.Println("*******不定参数输出*******")
	fmt.Println(join("hell", "o, wor", "ld", "!"))
	arr := []string{"hell", "o, wor", "ld", "!"}
	fmt.Println(join(arr...))
}

func getTwoNum(input int) (int, int) {
	where()
	return 1 * input, 2 * input

}

func getTriNum() (num1 int, num2 int, num3 int) {
	num1, num2, num3 = 100, 200, 300 // 可注释
	// ss := 4
	// ll := "ll"
	// return num2, num1, num3
	return

}

func getMap() (map[int]int, bool) {
	return nil, false
}

func getSlice() ([][]int, bool) {
	return nil, false
	// return [][]int{}, false
}

func join(arr ...string) string {
	for _, str := range arr[1:] {
		arr[0] += str
	}
	return arr[0]
}

/*
Conclusion:
	1. 可以先使用后定义
	2. 不可以在函数中定义函数, 除非立马返回
	3. 在一个函数中定义的数,在不传参的情况下, 不能给在它其中调用的函数用, 也就是参数作用域的问题
	4. 输出与输出均需要申明
	5. 输出的申明, 加上返回变量名, 可以不用在函数体中再定义, 并且, 返回只需要return
	6. 对于5, 如果要强行在return后加上变量名, go不会检查返回的顺序也不会检查返回变量名与实际是否一致, 只会检查变量个数与类型
	7. 对于5, 可以在函数体中只写return, 这样就获得了在申明中定义的返回值的初值
	8. TODO 变长参数
	9. TODO defer 与 finally
	10. nil可作为array, slice, map的返回值, 返回值会被自动当转化为对应的空对象, 不过是nil还是空的容器, 可以通过r==nil判断
	11. array, slice, map做参数传递时, 可以理解为传递的是指针
	12. ...非常类似于python中的*
*/
