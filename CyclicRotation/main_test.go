package main

import "testing"

func TestRotateSliceElements(t *testing.T) {

	type test struct {
		originalSlice []int
		shift         int
		answer        []int
	}

	tests := []test{
		test{[]int{1, 2, 3, 4, 5}, 1, []int{5, 1, 2, 3, 4}},
		test{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		test{[]int{1, 2, 3, 4, 5}, 3, []int{3, 4, 5, 1, 2}},
		test{[]int{1, 2, 3, 4, 5}, 4, []int{2, 3, 4, 5, 1}},
		test{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		test{[]int{1, 2, 3, 4, 5}, 6, []int{5, 1, 2, 3, 4}},
		test{[]int{0, 0, 0}, 1, []int{0, 0, 0}},
		test{[]int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
		test{[]int{3, 8, 9, 7, 6}, 3, []int{9, 7, 6, 3, 8}},
	}

	for _, v := range tests {
		X := rotateSliceElements(v.originalSlice, v.shift)

		for i := 0; i < len(X); i++ {
			if X[i] != v.answer[i] {
				t.Error("Expected: ", v.answer, "Got: ", X)
			}
		}

	}
}
