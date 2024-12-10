package main

import (
	"context"
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

func stopWithChannel() {
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Stopped with channel")
				return
			default:
				fmt.Println("Running...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	done <- true
	time.Sleep(1 * time.Second)
}

func stopWithContext() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Stopped with context")
				return
			default:
				fmt.Println("Running...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func stopWithTimeAfter() {
	stop := time.After(500 * time.Millisecond)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Stopped with time.After")
				return
			default:
				fmt.Println("Running...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)
}

func stopWithAtomicFlag() {
	var stopFlag atomic.Bool

	go func() {
		for {
			if stopFlag.Load() { // Проверяем флаг
				fmt.Println("Stopped with atomic flag")
				return
			}
			fmt.Println("Running...")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(1 * time.Second)
	stopFlag.Store(true)
	time.Sleep(1 * time.Second)
}

func main() {
	fmt.Println("Stopping with channel:")
	stopWithChannel()
	time.Sleep(1 * time.Second)

	fmt.Println("Stopping with context:")
	stopWithContext()
	time.Sleep(1 * time.Second)

	fmt.Println("Stopping with time.After:")
	stopWithTimeAfter()
	time.Sleep(1 * time.Second)

	fmt.Println("Stopping with atomic flag:")
	stopWithAtomicFlag()
	time.Sleep(1 * time.Second)

	fmt.Println("Stopping by killing main goroutine:")
	os.Exit(1)
}
