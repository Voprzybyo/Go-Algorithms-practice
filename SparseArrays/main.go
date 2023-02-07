package main

import "fmt"

func matchingStrings(strings []string, queries []string) []int {

	ret := make([]int, len(queries))

	for i, v1 := range queries {
		for _, v2 := range strings {
			if v1 == v2 {
				ret[i]++
			}
		}
	}
	return ret
}

func main() {
	strings := []string{"ab", "ab", "abc"}
	queries := []string{"ab", "abc", "bc"}

	ret := matchingStrings(strings, queries)
	fmt.Println(ret)

}
