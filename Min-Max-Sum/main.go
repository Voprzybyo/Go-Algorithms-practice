package main

import (
	"errors"
	"fmt"
	"sort"
)

func getMinMaxSum(N int, X []int) (int, int, error) {
	
	sumMin, sumMax := 0, 0
	sort.Ints(X)

	if len(X) < N {
		return 0, 0, errors.New("wrong slice size")
	}

	for i := 0; i < N; i++ {
		sumMin += X[i]
	}

	for i := len(X)-N; i < len(X); i++ {
		sumMax += X[i]
	}

	return sumMin, sumMax, nil
}

func main() {

	X := []int{1, 3, 2, 5, 4}
	N := 4

	min, max, err := getMinMaxSum(N, X)
	if err != nil {
		panic(err)
	}

	fmt.Println(min, max)
}
