package main

import "fmt"

func binarySearch(array []int, left, right, target int) int {
	if left <= right {
		mid := (right + left) / 2
		if array[mid] == target {
			return mid
		} else if array[mid] < target {
			return binarySearch(array, mid+1, right, target)
		} else {
			return binarySearch(array, left, mid-1, target)
		}
	} else {
		return -1
	}
}

func main() {
	var array = []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(binarySearch(array, 0, len(array)-1, 9))
}
