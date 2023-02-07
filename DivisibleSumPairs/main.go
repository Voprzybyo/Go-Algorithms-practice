package main

import "fmt"

func getNumberOfDivisiblePairs(k int, arr []int) int {
	
	counter := 0

	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			if (arr[i] + arr[j]) % k == 0 {
				counter++
			}
		}
	}
	
	return counter
}

func main() {

	k := 5
	X := []int{1, 2, 3, 4, 5, 6}

	numberOfPairs := getNumberOfDivisiblePairs(k, X)
	fmt.Println(numberOfPairs)
}
