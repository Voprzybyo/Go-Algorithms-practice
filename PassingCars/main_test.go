package main

import "testing"

func TestGetPassingCarsPairs(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{0, 1, 0, 1, 1}, 5},
		{[]int{0, 1, 1, 1, 1}, 4},
		{[]int{0, 1, 1, 0, 0}, 2},
	}

	for _, v := range tests {
		x := getPassingCarsPairs(v.data)
		if x != v.answer {
			t.Error("Expected: ", v.answer, "Got: ", x)
		}
	}
}
