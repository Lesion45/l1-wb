package main

import (
	"fmt"
	"sync"
)

func square(in <-chan int, out chan<- int) {
	defer close(out)
	for val := range in {
		square := val * val
		out <- square
	}
}

func read(out <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range out {
		fmt.Println(val)
	}
}

func main() {
	var wg sync.WaitGroup

	in := make(chan int, 10)
	out := make(chan int, 10)

	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	go square(in, out)
	wg.Add(1)
	go read(out, &wg)

	for _, val := range arr {
		in <- val
	}

	close(in)
	wg.Wait()
}
