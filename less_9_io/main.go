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
}
