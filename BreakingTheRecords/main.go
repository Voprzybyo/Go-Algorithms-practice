package main

import "fmt"

func getRecordBreakingCount(scores []int32) []int32 {

	var least, leastCounter int32 = scores[0], 0
	var most, mostCounter int32 = scores[0], 0

	for _, v := range scores {

		if v < least {
			least = v
			leastCounter++
		} else if v > most {
			most = v
			mostCounter++
		}
	}

	ret := []int32{mostCounter, leastCounter}
	return ret
}

func main() {
	scores := []int32{12, 24, 10, 24}
	ret := getRecordBreakingCount(scores)

	fmt.Println(ret)
}
