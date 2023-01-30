package main

import "testing"

func TestGetEarliestTime(t *testing.T) {

	type test struct {
		X      int
		A      []int
		answer int
	}

	tests := []test{
		{5, []int{1, 3, 1, 4, 2, 3, 5, 4}, 6},
		{3, []int{1, 4, 2, 3}, 3},
		{3, []int{1, 3}, -1},
		{4, []int{1, 2, 3, 2, 3, 3, 1, 2, 4, 4, 2, 1}, 8},
		{2, []int{1, 1, 1, 1}, -1},
		{4, []int{1, 2, 3, 4, 3, 3, 1, 2, 4, 4, 2, 1}, 3},
		{2, []int{1, 4, 2, 3}, 2},
		{4, []int{1, 2, 3, 2, 3, 3, 1, 2, 2, 4, 2, 1}, 9},
	}

	for _, v := range tests {
		x := getEarliestTime(v.X, v.A)
		if x != v.answer {
			t.Error("Got: ", x, "Expected: ", v.answer)
		}
	}
}
