package main

import (
	"fmt"
	"sort"
	"time"
)

func getMissingIntegerWithSorting(A []int) int {

	sort.Ints(A)
	min := 1

	for _, v := range A {
		if v == min {
			min++
		}
	}
	return min
}

func getMissingIntegerWithoutSorting(A []int) int {

	for i := 1; i < 100000; i++ {
		for j, v := range A {
			if v == i {
				break
			}
			if j == len(A)-1 {
				return i
			}
		}
	}
	return 1
}

func main() {
	A := []int{1, 2, 3, 4, 6, 7, 8, 9}

	start := time.Now()
	retWithoutSorting := getMissingIntegerWithoutSorting(A)
    elapsed1 := time.Since(start)

	start = time.Now()
	retWithSorting := getMissingIntegerWithSorting(A)
	elapsed2 := time.Since(start)
	
	fmt.Println("[WITHOUT SORTING] Missing integer: ", retWithoutSorting, "Execution time: ", elapsed1.Nanoseconds())
	fmt.Println("[WITH SORTING] Missing integer: ", retWithSorting, "Execution time: ", elapsed2.Nanoseconds())
}
