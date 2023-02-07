package main

import "testing"

func TestGetNumberOfDivisiblePairs(t *testing.T) {

	type test struct {
		k      int
		data   []int
		answer int
	}

	tests := []test{
		{5, []int{1, 2, 3, 4, 5, 6}, 3},
		{3, []int{1, 3, 2, 6, 1, 2}, 5},
	}

	for _, v := range tests {
		x := getNumberOfDivisiblePairs(v.k, v.data)

		if x != v.answer {
			t.Error("Got:", x, "Expected:", v.answer)
		}
	}
}
