package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]

	var less, equal, greater []int
	for _, num := range arr {
		if num < pivot {
			less = append(less, num)
		} else if num == pivot {
			equal = append(equal, num)
		} else {
			greater = append(greater, num)
		}
	}

	return append(append(quickSort(less), pivot), quickSort(greater)...)
}

func main() {
	arr := []int{10, 5, 2, 3, 8, 6, 1, 7, 4, 9}
	fmt.Println(quickSort(arr))
}
