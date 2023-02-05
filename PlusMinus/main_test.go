package main

import "testing"

func TestPlusMinus(t *testing.T) {

	type test struct {
		data   []int
		ansPos float32
		ansNeg float32
		ansZer float32
	}

	tests := []test{
		{[]int{1,2,3,4,5,-2,-3,0,0,0}, 0.5, 0.2, 0.3},
		{[]int{1,2,3,4,5,-2,-3,0,0},  0.5555556, 0.22222222, 0.22222222},
		{[]int{-1,-2,-3,-4,-5,-2,-3,0,0},  0, 0.7777778, 0.22222222},
		{[]int{0,0,0,0,0,0,0,0,0,0},  0, 0, 1},
	}

	for _, v := range tests {
		p, n, z := plusMinus(v.data) 
		if p != v.ansPos || n != v.ansNeg || z != v.ansZer {
			t.Error("Expected", v.ansPos, v.ansNeg, v.ansZer, "Got: ", p, n, z)
		}
	}
}
