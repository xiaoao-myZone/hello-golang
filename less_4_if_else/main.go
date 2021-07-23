package main

import (
	"fmt"
)

func main() {
	i, j := 1, 1
	if true {
		fmt.Println("lalla")
	} else if true { // else 不能换行
		fmt.Println("haala")
	} else {
		fmt.Println("waawa")
	}

	if cons := 5; cons < 2 {
		fmt.Println("yes")
	} else {
		fmt.Println("No")
	}

	if i = i + j; j > 0 {
		fmt.Println(i, j)
	}

	// fmt.Println(cons) cons is just valid in if - else
}

/*
Conclusion:
	1. if [expression;] bool expression {} [else]
*/
