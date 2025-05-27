package main

import "fmt"

func quicksort(array []int, left, right int) {
	if left < right {
		pivot := partition(array, left, right)
		quicksort(array, left, pivot-1)
		quicksort(array, pivot+1, right)
	}
}

func partition(array []int, left, right int) int {
	pivot := array[right]
	i := left - 1
	for j := left; j < right; j++ {
		if array[j] <= pivot { // moves j pointer until we find an element less or equal than the pivot
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i+1], array[right] = array[right], array[i+1]
	return i + 1
}

func main() {
	arr := []int{0, 3, 6, 7, 8, 4, 2, 1, 5}
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
