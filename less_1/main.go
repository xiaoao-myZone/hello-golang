package main

import (
	"./math_utils"
	"fmt"
)

func main() {
	var nums = 5
	var client = "agag"
	var ret = fmt.Sprintf("msg:%s//code:%d", client, my_lib.Rminus(nums, nums))
	nums = "555"
	fmt.Println(ret)

}

/*
Conclusion
	0. 貌似没有模块的概念, 导入只有包这个概念
	1. import 多个包用 import ("a" "b")
	2. ./dirName可以用来表示相对路径, 可不可以用../dirName?
	3. Sprintf可以格式化输出, 类似与python的 string.format(...)
	4. 定义变量, 但是不需要指定类型, 但是不像python那样呈现出多态, 一旦变量类型根据赋值确定, 便不能改变类型
	5. 字符串只能用双引号表示
*/
