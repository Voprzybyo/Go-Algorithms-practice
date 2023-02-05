package main

import "testing"

func TestGetMinMaxSum(t *testing.T) {
	type test struct {
		N      int
		data   []int
		ansMin int
		ansMax int
	}

	tests := []test{
		{4, []int{1, 3, 5, 7, 9}, 16, 24},
		{4, []int{1, 2, 3, 4, 5}, 10, 14},
		{3, []int{1, 3, 5, 7, 9}, 9, 21},
		{3, []int{9, 7, 5, 3, 1}, 9, 21},
		{2, []int{10, 10, 10, 10, 10}, 20, 20},
	}

	for _, v := range tests {
		resMin, resMax, err := getMinMaxSum(v.N, v.data)

		if err != nil{
			panic(err)
		}

		if resMin != v.ansMin || resMax != v.ansMax {
			t.Error("Expected: ", v.ansMin, v.ansMax, "Got: ", resMin, resMax)
		}
	}
}
