package main

import "fmt"

func singleNumber(nums []int) int {
	x := 0
	for _, n := range nums {
		x ^= n
	}
	return x
}

func main() {
	testcases := [][]int{
		{2, 2, 1},
		{4, 1, 2, 1, 2},
		{1},
	}
	for _, testcase := range testcases {
		fmt.Println(singleNumber(testcase))
	}
}
