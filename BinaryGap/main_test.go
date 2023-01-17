package main

import "testing"

func TestGetLongestBinaryGap(t *testing.T) {

	type test struct {
		data   int
		answer int
	}

	tests := []test{
		test{9, 2},
		test{529, 4},
		test{20, 1},
		test{15, 0},
		test{32, 0},
		test{1041, 5},
	}

	for _, v := range tests {
		x := getLongestBinaryGap(v.data)
		if x != v.answer {
			t.Error("Expected:", v.answer, "Got:", x)
		}
	}
}
