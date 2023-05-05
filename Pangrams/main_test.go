package main

import "testing"
func TestIsPangram(t *testing.T) {

	type test struct {
		inputStr string
		answer string
	}

	tests := []test {
		test{"We promptly judged antique ivory buckles for the next prize", "pangram"},
		test{"We promptly judged antique ivory buckles for the prize", "not pangram"},
		test{"the quick brown fox jumps over the lazy dog", "pangram"},
	}

	for _, v := range tests {
		x := isPangram(v.inputStr)
		if x != v.answer {
			t.Error("Expected:", v.answer, "Got:", x)
		}
	}

}