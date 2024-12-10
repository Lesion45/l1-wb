package main

import "fmt"

func removeElemFromSlice[T any](arr []T, index int) []T {
	return append(arr[:index], arr[index+1:]...)
}

func main() {
	numbers := []int{10, 20, 30, 40, 50, 90, 60}
	fmt.Println("Original Slice:", numbers)

	var index int
	fmt.Scan(&index)

	numbers = removeElemFromSlice(numbers, index)

	fmt.Println("Slice after deleting:", numbers)

}
