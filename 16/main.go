package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]

	var left, mid, right []int
	for _, num := range arr {
		if num < pivot {
			left = append(left, num)
		} else if num == pivot {
			mid = append(mid, num)
		} else {
			right = append(right, num)
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func main() {
	arr := []int{10, 5, 2, 3, 8, 6, 1, 7, 4, 9}
	fmt.Println(quickSort(arr))
}
