package main

import "fmt"

func marsExploration(s string) int32 {

	controlStr := "SOS"
	var changedLetters int32

	for i := 0; i < len(s); i++ {
		if s[i] != controlStr[i % 3] {
			changedLetters++
		}
	}

	return changedLetters
}

func main() {
	res := marsExploration("SOSSPSSQSSOR")
	fmt.Println(res)
}
