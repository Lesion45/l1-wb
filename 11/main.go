package main

import "fmt"

func main() {
	var set1 = make(map[int]struct{})
	var set2 = make(map[int]struct{})
	var intersection = make(map[int]interface{})

	for i := 1; i <= 10; i++ {
		set1[i] = struct{}{}
	}
	for i := 1; i <= 10; i++ {
		set2[i*3] = struct{}{}
	}

	for i := range set1 {
		if _, ok := set2[i]; ok {
			intersection[i] = struct{}{}
		}
	}

	fmt.Println("set1: ", set1)
	fmt.Println("set2: ", set2)
	fmt.Println("Пересечение set1 и set2: ", intersection)
}
