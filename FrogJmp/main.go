package main

import "fmt"

func MinimalNumberOfJumps(X int, Y int, D int) int {

	// Calculate distance to targer
	distanceToTarget := Y - X

	// Return number of necessary jumps
	if distanceToTarget%D == 0 {
		return distanceToTarget / D
	} else {
		return distanceToTarget/D + 1
	}
}

func main() {
	ret := MinimalNumberOfJumps(10, 85, 30)
	fmt.Println("Number of necessary jumps: ", ret)
}
