package main

import "testing"

func TestDiagonalDifference(t *testing.T) {

	type test struct {
		data   [][]int32
		answer int32
	}

	tests := []test {
		{[][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}, 2},
		{[][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}, 15},
		{[][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 8}}, 3},
		{[][]int32{{3, 2, 3}, {4, 5, 6}, {9, 8, 8}}, 1},
		{[][]int32{{1, 2, 3}, {4, 10, 6}, {9, 8, 8}}, 3},
	}

	for _, v := range tests {
		ret := getDiagonalDifference(v.data) 

		if ret != v.answer {
			t.Error("Got:", ret, "Expected:", v.answer)
		}
	}
}
