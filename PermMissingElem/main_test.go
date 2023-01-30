package main

import "testing"
//import "fmt"

func TestGetMissingElement(t *testing.T) {
	
	type test struct {
		data []int
		answer int
	}

	tests := []test {
		test{[]int{2,3,1,5}, 4},
		test{[]int{1,2,3,4,5,6,8,9,10}, 7},
		test{[]int{6,4,5,3,1,2,8,9,11,10}, 7},
		test{[]int{1,3,2,5,7,6,8,9}, 4},
		test{[]int{}, 1},
		test{[]int{1}, 2},
	}

	for _, v := range tests {
		x := getMissingElement(v.data)
		if x != v.answer {
			t.Error("Expected:", v.answer, "Got:", x)
		}
	}
}