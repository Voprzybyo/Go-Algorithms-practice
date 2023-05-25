package main

import "testing"

func TestMigratoryBird(t *testing.T) {
	type test struct {
		input  []int32
		answer int32
	}

	tests := []test{
		{[]int32{1, 1, 2, 2, 3}, 1},
		{[]int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}, 3},
	}

	for _, v := range tests {
		x := migratoryBirds(v.input)

		if x != v.answer {
			t.Error("Expected:", v.answer, "Got:", x)
		}
	}
}
