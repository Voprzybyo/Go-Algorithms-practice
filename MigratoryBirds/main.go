package main

import "fmt"

func migratoryBirds(birds []int32) int32 {

	m := make(map[int32]int32)
	var max int32 = 0
	var lowestId int32 = 1
	
	for _, v := range birds {
		m[v]++
	}

	for i, v := range m {
		if v > max {
			max = v
			lowestId = i
		}
	}

	return lowestId
}

func main() {

	birds := []int32{1, 1, 2, 2, 3}
	lowestId := migratoryBirds(birds)

	fmt.Println(lowestId)
}
