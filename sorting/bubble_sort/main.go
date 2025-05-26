package main

import "fmt"

func bubble(nums []int) {
	size := len(nums)
	for range nums {
		isSorted := true
		for i := 0; i < (size - 1); i++ {
			if nums[i] > nums[i+1] {
				isSorted = false
				nums[i+1], nums[i] = nums[i], nums[i+1]
			}
		}
		if isSorted {
			return
		}
	}
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	bubble(arr)
	fmt.Println(arr)
}
