package main

import "fmt"

func getEarliestTime(X int, A []int) int {

	if X > len(A) {
		return -1
	}

	lackingLeafs := X
	m := make(map[int]bool)

	for i, v := range A {

		if v > X {
			continue
		}

		if !m[v] {
			m[v] = true // Position is covered
			lackingLeafs -= 1
			if lackingLeafs == 0 {
				return i
			}
		}
	}
	return -1
}

func main() {

	X := 2
	A := []int{1, 4, 2, 3}
	ret := getEarliestTime(X, A)
	fmt.Println("Earliest time: ", ret)
}
