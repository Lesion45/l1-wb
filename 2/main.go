package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	results := make(chan int, len(numbers))

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			results <- n * n
		}(num)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for square := range results {
		fmt.Println(square)
	}
}
