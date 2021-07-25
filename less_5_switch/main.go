package main

import (
	"fmt"
)

func main() {
	var t int = 56
	fmt.Println(t)
	// case-1
	switch t {
	case 55:
		{

			fmt.Println("this is 55")
		}
	case 56:
		fallthrough
		// {
		// 	fmt.Println("this is 56")
		// 	fmt.Println("this is *56*")
		// 	// fallthrough
		// }
		// return

	case 57:
		fmt.Println("this is 57")
	default:
		fmt.Println("this is a num above 57")
	}

	// case-2
	switch {
	case t <= 100:
		fmt.Println("less than 100")
	case t <= 200:
		fmt.Println("less than 200")
	default:
		fmt.Println("greater than 200")

	}

}

/*
Conclusion:
	1. switch 后的参数是否只能整型?
	2. 命中了一个case后, 是否会执行其他case
	3. 如何理解fallthrough, 貌似只能作为一个case的唯一一行语句 TODO
	4. 但是fallthrough可以在不写{}是加入到case之间,如果加在命中的case后, 会使程序无脑执行下一个case
*/
