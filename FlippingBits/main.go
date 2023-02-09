package main

import (
	"fmt"
	"math"
)

func flipBits(num int64) int64 {

	size := math.Pow(2, 32)
	ret := num ^ (int64(size) - 1)

	return ret
}

func main() {

	var input int64 = 2147483647
	ret := flipBits(input)
	fmt.Println(ret)
}
