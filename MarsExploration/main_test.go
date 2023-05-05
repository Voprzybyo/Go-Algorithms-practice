package main

import "testing"

func TestMarsExploration(t *testing.T) {
	type test struct {
		inputStr string
		answer   int32
	}

	tests := []test {
		test{"SOSSPSSQSSOR", 3},
		test{"SOSSOT", 1},
		test{"SOSSOSSOS", 0},
	}

	for _, v := range tests {
		x := marsExploration(v.inputStr)
		if x != v.answer {
			t.Error("Expected:", v.answer, "Got:", x)
		}
	}
}
