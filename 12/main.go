package main

import "fmt"

func main() {
	var strs = []string{"cat", "cat", "dog", "cat", "tree"}
	var set = make(map[string]struct{}, len(strs))

	for _, el := range strs {
		set[el] = struct{}{}
	}

	fmt.Println(set)
}
