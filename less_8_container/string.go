package main

import "fmt"

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
}
