package main

import "testing"

func TestFlipBits(t *testing.T) {

	type test struct {
		input  int64
		output int64
	}

	tests := []test {
		{4, 4294967291},
		{123456, 4294843839},
		{0, 4294967295},
		{802743475, 3492223820},
		{35601423, 4259365872},
		{2147483647, 2147483648},
		{1, 4294967294},
	}

	for _, v := range tests {
		x := flipBits(v.input)

		if x != v.output {
			t.Error("Expected:", v.output, "Got:", x)
		}
	}
}
