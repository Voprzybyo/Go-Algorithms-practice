package main

import "testing"

func TestCheckPermutation(t *testing.T) {
	type test struct {
		data   []int
		answer bool
	}

	tests := []test{
		{[]int{4, 1, 3, 2}, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 10}, false},
		{[]int{1}, true},
		{[]int{1, 3, 2, 4, 5, 6, 8, 7, 10, 9}, true},
		{[]int{1, 3, 2, 4, 5, 6, 8, 7, 9}, true},
	}

	for _, v := range tests {
		x := checkPermutation(v.data)
		if x != v.answer {
			t.Error("Got: ", x, "Expected: ", v.answer)
		}
	}
}
