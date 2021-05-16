package main

import "fmt"

func main() {
	s := "hello, world!"
	c := []byte(s)
	fmt.Printf("%#v\n%+v\n", c, c)
}
