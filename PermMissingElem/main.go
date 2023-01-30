package main

import "fmt"

func getMissingElement(A []int) int {
	
	// Sum all values in slice without missing element
	sumWithoutMissing := 0 
	for i := 1; i <= len(A) + 1; i++ {
		sumWithoutMissing += i
	}

	// Sum all values in slice with missing element
	sumWithMissing := 0
	for _, v := range A {
		sumWithMissing += v
	}

	return sumWithoutMissing - sumWithMissing
}

func main() {
	
	A := []int{2,3,1,5}

	ret := getMissingElement(A)
	fmt.Println("Missing element: ", ret)
}
