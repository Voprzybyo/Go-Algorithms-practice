package main

import "testing"

func TestToundGrades(t *testing.T) {

	type test struct {
		data   []int32
		answer []int32
	}
	
	tests := []test{
		{[]int32{73, 67, 38, 33}, []int32{75, 67, 40, 33}},
		{[]int32{72, 68, 38, 33, 39}, []int32{72, 70, 40, 33, 40}},
		{[]int32{1, 2, 3, 4, 5, 6, 7}, []int32{1, 2, 3, 4, 5, 6, 7}},
		{[]int32{38, 39, 40, 41, 42, 43}, []int32{40, 40, 40, 41, 42, 45}},
		{[]int32{44, 45, 46, 47, 48}, []int32{45, 45, 46, 47, 50}},
	}

	for _, v1 := range tests {
		roundedGrades := roundGrades(v1.data)

		for i, v2 := range roundedGrades {
			if v2 != v1.answer[i] {
				t.Error("Expected:", v1.answer[i], "Got:", v2)
			}
			
		}
	}
}
