package main

import "fmt"

func binarySearch(nums []int, key int) int {
	low, high := 0, len(nums)
	for low < high {
		mid := (low + high) / 2
		if nums[mid] == key {
			return mid
		}
		if nums[mid] < key {
			low = mid + 1
		}
		if nums[mid] > key {
			high = mid
		}
	}
	return -1
}

func main() {
	inputs := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
	}
	for _, input := range inputs {
		fmt.Println(binarySearch(input, 3))
	}
}
