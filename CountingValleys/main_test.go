package main

import "testing"

func TestCountingValleys(t *testing.T) {

	type test struct {
		steps  int32
		path   string
		answer int32
	}

	tests := []test{
		test{8, "UDDDUDUU", 1},
		test{12, "DDUUDDUDUUUD", 2},
	}

	for _, v := range tests {
		x := countingValleys(v.steps, v.path)
		if x != v.answer {
			t.Error("Expected:", v.answer, "Got:", x)
		}
	}

}
