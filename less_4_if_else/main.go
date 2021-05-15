package main

import (
	"fmt"
)

func main() {
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

	// fmt.Println(cons) cons is just valid in if - else
}
