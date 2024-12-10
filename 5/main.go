package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Write 1 arg: time of work in s")
		return
	}
	t, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("You need to write a number")
		return
	}

	ch := make(chan interface{}, 1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(t)*time.Second)
	var wg sync.WaitGroup
	defer cancel()

	wg.Add(2)
	go func() {
		defer wg.Done()
		defer close(ch)

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Producer: Time out: ", ctx.Err())
				return
			default:
				ch <- rand.Int()
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Receiver: Time out: ", ctx.Err())
				return
			case data, ok := <-ch:
				if !ok {
					fmt.Println("Channel closed")
					return
				}
				fmt.Println("Data received:", data)
			}
		}
	}()

	wg.Wait()
}
