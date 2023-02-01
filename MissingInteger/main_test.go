package main

import "testing"

func TestGetMissingInteger(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{1, 2, 3, 4, 6, 7, 8, 9}, 5},
		{[]int{3, 4, 1, 5, 5, 7, 6, 9}, 2},
		{[]int{1, 2, 3}, 4},
		{[]int{-1, -3}, 1},
	}

	for _, v := range tests {
		x := getMissingIntegerWithoutSorting(v.data) 
		if x != v.answer {
			t.Error("[Without sorting] Got: ", x, "Expected: ", v.answer)
		}
	}

	for _, v := range tests {
		x := getMissingIntegerWithSorting(v.data) 
		if x != v.answer {
			t.Error("[With sorting] Got: ", x, "Expected: ", v.answer)
		}
	}
}
