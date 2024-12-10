package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	left, right := 0, len(words)-1

	for left <= right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}
	fmt.Println(words)

	return strings.Join(words, " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Something went wrong: ", err)
	}
	s = strings.TrimSpace(s)
	fmt.Println(reverseWords(s))
}
