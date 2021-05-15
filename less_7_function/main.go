package main

import "fmt"

var num int = 10

func main() {
	var k, v, s int
	k, v = getTwoNum(num)
	k, v, s = getTriNum()
	fmt.Println(k, v, s)
}

func getTwoNum(input int) (int, int) {
	return 1 * input, 2 * input
}

func getTriNum() (num1 int, num2 int, num3 int) {
	num1, num2, num3 = 100, 200, 300 // 可注释
	// ss := 4
	// ll := "ll"
	// return num2, num1, num3
	return

}

/*
Conclusion:
	1.可以先使用后定义
	2.不可以在函数中定义函数
	3. 在一个函数中定义的数,在不传参的情况下, 不能给在它其中调用的函数用, 也就是参数作用域的问题
	4. 输出与输出均需要申明
	5. 输出的申明, 加上返回变量名, 可以不用在函数体中再定义, 并且, 返回只需要return
	6. 对于5, 如果要强行在return后加上变量名, go不会检查返回的顺序也不会检查返回变量名与实际是否一致, 只会检查变量个数与类型
	7. 对于5, 可以在函数体中只写return, 这样就获得了在申明中定义的返回值的初值
*/
