package main

import (
	"fmt"
	"math"
)

func binarySearch(nums []int, key int, low int, high int) int {
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

func exponentialSearch(nums []int, key int) int {
	if nums[0] == key {
		return 0
	}
	n := len(nums)
	r := 1
	for r < n && nums[r] < key {
		r *= 2
	}
	if nums[r] == key {
		return r
	}
	return binarySearch(nums, key, r/2, int(math.Min(float64(r), float64(n-1))))
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	target := 32
	fmt.Printf("Element found at index: %d\n", exponentialSearch(nums, target))
}
