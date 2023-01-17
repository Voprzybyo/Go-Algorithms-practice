package main

import "fmt"

func rotateSliceElements(A []int, Shift int) []int {

	tempSlice := make([]int, len(A))
	copy(tempSlice, A)

	for i := 0; i < len(A); i++ {
		A[(i + Shift) % len(A)] = tempSlice[i % len(A)]
	}

	return A
}

func main() {
	
	A := []int{1, 2, 3, 4, 5, 6}
	Shift := 2

	// Print original slice of INTs
	// ============================
	fmt.Printf("Given shift:   %d\nGiven slice:   ", Shift)
	for _, v := range A {
		fmt.Printf("%d ", v)
	}

	ret := rotateSliceElements(A, Shift)

	// Print shifted slice of INTs
	// ===========================
	fmt.Printf("\nShifted slice: ")
	for _, v := range ret {
		fmt.Printf("%d ", v)
	}
}
