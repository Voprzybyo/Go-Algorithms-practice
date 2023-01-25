package main

import "fmt"

func getOddFromSlice(A []int) int {
	xorTempVal := 0
	for _, v := range A {
		xorTempVal ^= v
	}
	return xorTempVal
}

func main() {

	A := []int{9, 3, 9, 3, 9, 7, 9}
	ret := getOddFromSlice(A)
	fmt.Println(ret)
}
