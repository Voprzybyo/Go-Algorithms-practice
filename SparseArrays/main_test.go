package main

import "testing"

func TestMatchingStrings(t *testing.T) {

	type test struct {
		strings []string
		queries []string
		answer  []int
	}

	tests := []test{
		{[]string{"aba", "baba", "aba", "xzxb"}, []string{"aba", "xzxb", "ab"}, []int{2, 1, 0}},
		{[]string{"def", "de", "fgh"}, []string{"de", "lmn", "fgh"}, []int{1, 0, 1}},
		{[]string{"abcde", "sdaklfj", "asdjf", "na", "basdn", "sdaklfj", "asdjf", "na", "asdjf", "na", "basdn", "sdaklfj", "asdjf"},
			[]string{"abcde", "sdaklfj", "asdjf", "na", "basdn"}, []int{1, 3, 4, 3, 2}},
	}

	for _, v1 := range tests {

		x := matchingStrings(v1.strings, v1.queries)

		for i, v2 := range v1.answer {
			if x[i] != v2 {
				t.Error("Expected:", v2, "Got:", x[i])
			}
		}
	}
}
