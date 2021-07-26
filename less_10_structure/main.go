package main

import "fmt"

type s struct {
	arr []int
	ok  bool
}

func main() {
	m := s{arr: []int{}, ok: false}
	n := s{[]int{1, 2}, true}
	fmt.Println(m)
	fmt.Println(n)
}
