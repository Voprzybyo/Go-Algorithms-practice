package main

import "testing"

func TestGetRecordBreakingCount(t *testing.T) {

	type test struct {
		data   []int32
		answer []int32
	}

	tests := []test {
		{[]int32{12,24,10,24}, []int32{1,1}},
		{[]int32{3,4,21,36,10,28,35,5,24,42}, []int32{4,0}},
		{[]int32{10, 5, 20, 20, 4, 5, 2, 25, 1}, []int32{2,4}},
	}

	for _, v := range tests {
		x := getRecordBreakingCount(v.data)

		if x[0] != v.answer[0] || x[1] != v.answer[1] {
			t.Error("Got:", x, "Expected:", v.answer)
		}
	}
}
