package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data []int
		ans  int
	}
	tests := []test{
		{[]int{21, 21}, 42},
		{[]int{5, 8}, 13},
		{[]int{17, 30, 15}, 62},
		{[]int{-1, 0, 1}, 0},
	}

	for _, v := range tests {
		x := mySum(v.data...)

		if x != v.ans {
			t.Errorf("Expected: %d\nGot: %d", v.ans, x)
		}
	}

}
