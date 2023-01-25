package main

import "testing"
func TestMinimalNumberOfJumps(t *testing.T) {

	type test struct {
		data1   int
		data2   int
		data3   int
		answer int
	}

tests := []test {
	{10, 85, 30, 3},
	{10, 20, 5, 2},
	{5, 6, 1, 1},
	{5,6, 80, 1},
}

	for _, v := range tests {
		x := MinimalNumberOfJumps(v.data1, v.data2, v.data3)
		if(x != v.answer) {
			t.Error("Got: ", x, "Expected: ", v.answer)
		}
	}
}


