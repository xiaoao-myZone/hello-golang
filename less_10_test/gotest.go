package gotest

import (
	"fmt"
	"math"
)

func divide1(dividend int, divisor int) int {
	bits, res, isNegtive := 31, 0, false
	if dividend >= 0 && divisor >= 0 || dividend <= 0 && divisor <= 0 {
		isNegtive = true
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	fmt.Println(dividend, divisor, isNegtive)
	for {
		if dividend < divisor {
			break
		}
		for {
			if dividend>>bits >= divisor {
				break
			}
			bits--
			// fmt.Println(bits)

		}
		res += 1 << bits
		// fmt.Println(res)
		dividend = dividend - divisor<<bits

	}
	if isNegtive {
		if res > math.MaxInt32 {
			return math.MaxInt32
		} else {
			return res
		}

	} else {
		return -res
	}

}
