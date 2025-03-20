package main

import (
	"fmt"
	"math"
)

func maximumLengthSubstring(s string) int {
	l, r := 0, 0
	maxSize := 1
	counter := map[uint8]int{}
	counter[s[0]] = 1
	for r < len(s)-1 {
		r++
		if counter[s[r]] != 0 {
			counter[s[r]]++
		} else {
			counter[s[r]] = 1
		}
		for counter[s[r]] == 3 {
			counter[s[l]]--
			l++
		}
		maxSize = int(math.Max(float64(maxSize), float64(r-l+1)))
	}
	return maxSize
}

func main() {
	testcases := []string{"bcbbbcba", "aaaa"}
	for _, testcase := range testcases {
		fmt.Println(maximumLengthSubstring(testcase))
	}
}
