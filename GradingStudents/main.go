package main

import "fmt"

func roundGrades(grades []int32) []int32 {

	var roundedGrades []int32

	for _, v := range grades {
		mod := 5 - (v % 5)
		if v < 38 || mod > 2 {
			roundedGrades = append(roundedGrades, v)
		} else {
			roundedGrades = append(roundedGrades, v+mod)
		}
	}

	return roundedGrades
}

func main() {

	grades := []int32{73, 67, 38, 33}

	roundedGrades := roundGrades(grades)
	fmt.Println(roundedGrades)

}
