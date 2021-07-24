package main

import "fmt"

func main() {
	fmt.Println("hello, wolrd!")
	fmt.Println(5 / 2)
	a, b := 0, 1
	a, b = b, a
	fmt.Println(a, b)
	s := "1254a"
	fmt.Println(s[0] > 'a')
	fmt.Printf("%v\n", s[0])
}

/*
Conclusion:
	1. go run 只能运行 package 为 main 的包
	2. fmt 类似于 C中的 stdio
	3. 一行代码后面不需要加;
	4. 函数不需要声明返回值
	5. go build hello_world.go可以编译
	6. go的文件可以用'-'
	7. '{'不可以单独放在一行, 显然这不是在编译上有什么技术困难, 而是golang的作者对单行写'{'的方式深恶痛绝

*/
