package main

import (
	"fmt"
	"strings"
)

func isPangram(s string) string {

	isLetterPresent := make(map[int32]bool)
	s = strings.ToLower(s)
	
	for i := 'a'; i <= 'z'; i++ {
		isLetterPresent[i] = false
	}

	for i := 0; i < len(s); i++ {
		isLetterPresent[int32(s[i])] = true
	}

	for _, v := range isLetterPresent {
		if !v {
			return "not pangram"
		}
	}
	return "pangram"
}

func main() {

	res := isPangram("We promptly judged antique ivory buckles for the next prize")
	fmt.Println(res)
}
