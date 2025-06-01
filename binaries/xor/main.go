package main

import "fmt"

func missingNumber(nums []int) int {
	x := 0
	for _, n := range nums {
		x ^= n
	}
	for i := 0; i < len(nums)+1; i++ {
		x ^= i
	}
	return x
}

func main() {
	testcases := [][]int{
		{3, 0, 1},
		{0, 1},
		{9, 6, 4, 2, 3, 5, 7, 0, 1},
	}
	for _, testcase := range testcases {
		fmt.Println(missingNumber(testcase))
	}
}
