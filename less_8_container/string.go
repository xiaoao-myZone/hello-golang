package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello, world!"
	c := []byte(s)
	fmt.Printf("%#v\n%+v\n", c, c)
	fmt.Printf("%#v\n%+v\n", s[0], s[0])
	// 分片
	p := s[1:]
	fmt.Println(p)
	if s[2:4] == "ll" {
		fmt.Println("yes")
	} else {
		fmt.Println("No")
	}
	// 拼接
	s += "l"
	fmt.Println(s)

	// strings
	fmt.Println(strings.Contains(s, "l")) //string string bool
	fmt.Println(strings.Index(s, "x"))    //string string int
	fmt.Println(strings.Join([]string{"a", "b", "c"}, "+"))

}

/*
总结:
1. string可以用==来比较, 像数列一样, 可以索引, 可以分片
2. string可以像整数一样用+拼接, 这个和python类似
3. 更多方法需要引入strings这个内置包
	1. Contains 类似 python中的 in
	2. Join 与python用法一致
	3. Index 相当于find[python], 找不到返回-1
	4. Replace 相当于replace[python]
	5. Repeat 相当于count[python]
	6. Split 相当于split[python]
参考:
1. https://blog.csdn.net/cabing2005/article/details/60586189
*/
