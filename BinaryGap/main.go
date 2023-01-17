package main

import (
	"fmt"
	"strconv"
)

func getLongestBinaryGap(N int) int {

	// Change number to its binary representation (string)
	// ===================================================
	binaryNumber := strconv.FormatInt(int64(N), 2)
	fmt.Printf("Number: %d -> Binary representation: %s\n", N, binaryNumber)

	// Iterate through the string and calculate length of binary gap
	// =============================================================
	maxBinaryGapSize := 0
	binaryGapSize := 0

	for i := 1; i < len(binaryNumber); i++ {
		binaryGapSize++
		if binaryNumber[i] == '1' {
			if binaryGapSize > maxBinaryGapSize {
				maxBinaryGapSize = binaryGapSize - 1
			}
			binaryGapSize = 0
		}
	}
	return maxBinaryGapSize
}

func main() {

	const Number = 1041
	ret := getLongestBinaryGap(Number)

	fmt.Printf("Length of the longest binary gap in number %d: %d", Number, ret)
}
