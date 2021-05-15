package main

import "fmt"

const (
	x int = 5 + iota
	y
	z
)

func main() {
	var i = 5
	var j int = 4
	var p *int = nil             // nil竟是一个关键字
	fmt.Printf("%d, %p\n", i, p) // 居然是 0x0
	fmt.Println(i, j, p, x, y, z)
	fmt.Printf("%+v\n", p) // %+v输出 0xc0000140d8 %#v 输出 (*int)(0xc0000140d8)
	fmt.Println(*p)
}
