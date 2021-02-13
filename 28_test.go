package leetcode

import "testing"

func Test28(t *testing.T) {
	var fn = func(a, b string, ans int) {
		if i := strStr(a, b); i != ans {
			t.Fatal("失败", i)
		}
		t.Log("成功")
	}

	fn("", "", 0)
	fn("hello", "ll", 2)
	fn("golang", "ll", -1)
	fn("abc", "c", 2)
}
