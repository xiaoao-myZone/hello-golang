package main

import (
	"fmt"
	"os"
)

func main() {
	s := "hello, world!"
	d := 100
	ss := fmt.Sprintf("%#v\t%#v", s, d)
	ll := "\"af\""
	fmt.Println(s, d)
	fmt.Println(ss)
	fmt.Println(ll)
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
	fmt.Println(1 != 0 && 2 != 0)
	var r float64
	r = float64(d / 3.0)
	fmt.Printf("%f %f \n", r, 1000.0/d)
}

/*
Conclusion:
	1. 除法运算中有整数变量不可以与浮点变量运算
	2. 整数变量除以浮点常熟, 结果仍然是整数, 反之亦然
*/
