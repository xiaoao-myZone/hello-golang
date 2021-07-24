package gotest

import "testing"

type pairs struct {
	dividend int
	divisor  int
	res      int
}

func Test_Division(t *testing.T) {
	test_case := []pairs{
		{
			dividend: 1,
			divisor:  2,
			res:      0,
		},
		{
			10,
			2,
			5,
		},
		{
			100,
			11,
			9,
		},
		{
			-400,
			101,
			-3,
		},
		{
			2147483647,
			1,
			2147483647,
		},
		{
			-2147483648,
			-1,
			2147483647,
		},
	}
	for _, pair := range test_case {
		if i := divide1(pair.dividend, pair.divisor); i != pair.res {
			t.Error("未通过")
		} else {
			t.Log("通过")
		}
	}

}
