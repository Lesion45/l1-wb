package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	s = strings.ToLower(s)
	set := make(map[rune]struct{}, len(s))

	for _, r := range s {
		if _, ok := set[r]; ok {
			fmt.Println("false")
			return
		} else {
			set[r] = struct{}{}
		}
	}

	fmt.Println("true")
}
