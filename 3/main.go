package main

import (
	"fmt"
)

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	sum := 0

	ch := make(chan int, len(numbers))

	go func() {
		for i := 0; i < len(numbers); i++ {
			go func(num int) {

				square := num * num
				ch <- square
			}(numbers[i])
		}
	}()

	for i := 0; i < len(numbers); i++ {
		sum += <-ch
	}
	close(ch)
	fmt.Println(sum)
}
