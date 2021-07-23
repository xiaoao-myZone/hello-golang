package gotest

import "testing"

func Test_Division(t *testing.T) {
	if i := divide1(7, 1); i != 7 {
		t.Error("未通过")
	} else {
		t.Log("通过")
	}
}
