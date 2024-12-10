package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func reverseString(s string) string {
	runes := []rune(s)
	left, right := 0, len(runes)-1

	for left <= right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Something went wrong: ", err)
	}
	s = strings.TrimSpace(s)
	fmt.Println(reverseString(s))
}
