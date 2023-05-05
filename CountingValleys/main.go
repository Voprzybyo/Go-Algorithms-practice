package main

import "fmt"

func countingValleys(steps int32, path string) int32 {

	var i, counter, howManyValleys int32

	for i = 0; i < steps; i++ {
		if path[i] == 'U' {
			counter++
		} else if path[i] == 'D' {
			counter--
		}

		if counter == 0 && path[i] == 'U' {
			howManyValleys++
		}
	}
	return howManyValleys
}

func main() {

	res := countingValleys(8, "UDDDUDUU")
	fmt.Println(res)
}
