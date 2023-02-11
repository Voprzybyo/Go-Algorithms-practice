package main

import (
	"fmt"
	"math"
)

func getDiagonalDifference(X [][]int32) int32 {

	var sum1 int32 = 0
	for i := 0; i < len(X); i++ {
		for j := 0; j < len(X); j++ {
			if i == j {
				sum1 += X[i][j]
			}
		}
	}

	var sum2 int32 = 0
	for i := 0; i < len(X); i++ {
		for j := len(X); j >= 0; j-- {
			if i == len(X)-j-1 {
				sum2 += X[i][j]
			}
		}
	}

	return int32(math.Abs(float64(sum1) - float64(sum2)))
}

func main() {

	X := [][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}

	ret := getDiagonalDifference(X)
	fmt.Println(ret)

}
