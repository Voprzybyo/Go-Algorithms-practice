package main

import "testing"

func TestSockMerchant(t *testing.T) {
	type test struct {
		input  []int32
		length int32
		ans    int32
	}

	tests := []test{
		{[]int32{10, 20, 20, 10, 10, 30, 50, 10, 20}, 9, 3},
		{[]int32{1, 1, 3, 1, 2, 1, 3, 3, 3, 3}, 10, 4},
		{[]int32{1, 1, 2, 2, 3, 3}, 6, 3},
	}

	for _, v := range tests {
		x := sockMerchant(v.length, v.input)
		if x != v.ans {
			t.Error("Got:", x, "Expected:", v.ans)
		}
	}
}
