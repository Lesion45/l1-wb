package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var counter int

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}()

	wg.Wait()
	fmt.Println("Значение счетчика: ", counter)
}
