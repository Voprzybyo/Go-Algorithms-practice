package main

import "fmt"

func plusMinus(X []int) (float32, float32, float32) {

	var p, n, z float32 = 0.0, 0.0, 0.0
	sliceSizeI := len(X)
	var sliceSizeF float32 = float32(sliceSizeI)

	for _, v := range X {
		if v > 0 {
			p++
		} else if v < 0 {
			n++
		} else {
			z++
		}
	}

	return p / sliceSizeF, n / sliceSizeF, z / sliceSizeF
}

func main() {

	X := []int{-4, 3, -9, 0, 4, 1}
	positive, negative, zeros := plusMinus(X)

	fmt.Println(positive, negative, zeros)
}
