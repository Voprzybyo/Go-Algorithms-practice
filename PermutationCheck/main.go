package main

import "fmt"

func checkPermutation(A []int) bool {

	// Create and initialize map with "false" values
	m := make(map[int]bool)
	for i := range m {
		m[i] = false
	}

	for _, v := range A {
		
		if v > len(A) || v < 1 || m[v] {
			return false
		}
		m[v] = true
	}

	// Check if there is any false value in map
	for _, v := range m {
		if !v {
		return false
		}

	}
	return true
}

func main() {

	A := []int{1, 3, 2, 4, 5, 6, 8, 7, 9}
	
	isPermutation := checkPermutation(A)
	fmt.Println("Is given slice permutation? ->", isPermutation)
}
