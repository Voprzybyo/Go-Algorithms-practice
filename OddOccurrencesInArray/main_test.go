package main

import "testing"

func TestGetOddFromSlice(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{9, 3, 9, 3, 9, 7, 9}, 7},
		{[]int{1, 2, 1, 2, 3, 4, 4}, 3},
		{[]int{1, 2, 3, 3, 1}, 2},
		{[]int{1, 2, 1, 2, 1, 2, 1, 2, 1}, 1},
	}

	for _, v := range tests {

		x := getOddFromSlice(v.data)
		if x != v.answer {
			t.Error("Expected:", v.answer, "Got:", x)
		}
	}
}
