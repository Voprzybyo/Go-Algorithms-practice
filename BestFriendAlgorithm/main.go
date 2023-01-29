package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fillSliceWithFriends(ss [][]string) {

	rand.Seed(time.Now().UnixNano())
	names := []string{"John", "Paul", "Rupert", "Christopher", "Steven", "Mark", "James", "Anthony", "Hannah"}

	fmt.Printf("Randomized groups:\n")
	for i := 0; i < cap(ss); i++ {
		ss[i] = make([]string, 3)
		for j := 0; j < cap(ss[i]); j++ {
			randNumber := rand.Intn(len(names))
			ss[i][j] = names[randNumber]
		}
	}
	fmt.Println(ss)
	fmt.Println("")
}

func getBestFriends(ss [][]string) map[string][]string {

	friendship := make(map[string][]string)
	for i := 0; i < len(ss); i++ {
		for j := 0; j < len(ss[i]); j++ {
			for k := 0; k < 3; k++ {
				friendship[ss[i][j]] = append(friendship[ss[i][j]], ss[i][k])
			}
		}
	}
	return friendship
}

func removeElement(s []string, i int) []string {

	s[i] = ""
	return s[1:]
}

func CountOccurence(s []string) map[string]int {

	m := make(map[string]int)
	for _, v := range s {
		if v == "" {
			continue
		}
		m[v]++
	}
	return m
}

func pickBestFriends(m map[string][]string) {

	// Remove key string from map values
	for k, v := range m {
		for i, j := range v {
			if k == j {
				removeElement(v, i)
			}
		}
	}

	// Count occurences and check if sb occurs twice
	for k, v := range m {
		moreThanOne := false
		ret := CountOccurence(v)

		fmt.Printf("%s  :  { ", k)
		for _, j := range ret {
			if j == 2 {
				moreThanOne = true
			}
		}

		// Print group of best friends from data set
		if moreThanOne {
			for i, j := range ret {
				if j == 2 {
					fmt.Printf("%s  ", i)
				}
			}
		} else {
			for i, j := range ret {
				if j == 1 {
					fmt.Printf("%s  ", i)
				}
			}
		}
		fmt.Println("}")
	}
}

func main() {

	Input := make([][]string, 3)

	fillSliceWithFriends(Input)
	allFriends := getBestFriends(Input)
	pickBestFriends(allFriends)
}
