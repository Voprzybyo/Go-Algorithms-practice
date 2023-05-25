package main

import "fmt"

func sockMerchant(n int32, ar []int32) int32 {

	m := make(map[int32]int32)
	var counter int32

	for _, v := range ar {
		m[v]++
	}

	for _, v := range m {
		if v%2 == 0 {
			counter += v / 2
		} else {
			counter += (v - 1) / 2
		}
	}
	return counter
}

func main() {

	ar := []int32{1, 1, 3, 1, 2, 1, 3, 3, 3, 3}

	res := sockMerchant(10, ar)

	fmt.Println(res)
}
