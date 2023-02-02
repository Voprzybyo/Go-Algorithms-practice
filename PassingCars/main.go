package main

import "fmt"

func getPassingCarsPairs(A []int) int {

	eastCar, ret := 0, 0
	
	for i := range A {

		if A[i] == 0 {
			eastCar += 1
		} else {
			ret += eastCar
		}

		if ret > 1000000000 {
			return -1
		}
	}
	return ret
}

func main() {

	A := []int{0, 1, 0, 1, 1}

	retPairs := getPassingCarsPairs(A)
	fmt.Println("Number of pairs: ", retPairs)
}
