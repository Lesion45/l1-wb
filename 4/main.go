package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type Worker struct {
	id int
	wg *sync.WaitGroup
	ch chan interface{}
}

func NewWorker(id int, wg *sync.WaitGroup, ch chan interface{}) *Worker {
	return &Worker{
		id: id,
		wg: wg,
		ch: ch,
	}
}

func (w *Worker) Run(ctx context.Context) {
	defer w.wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context canceled. Worker id: ", w.id)
			return
		case data, ok := <-w.ch:
			if !ok {
				fmt.Println("Channel closed. Worker id: ", w.id)
				return
			}
			fmt.Println("Data received: ", data)
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Write 1 arg: count of workers")
		return
	}
	k, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("You need to write a number")
		return
	}

	var wg sync.WaitGroup

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan interface{}, k)

	for i := 0; i < k; i++ {
		go func(id int) {
			wg.Add(1)
			worker := NewWorker(id, &wg, ch)
			worker.Run(ctx)
		}(i)
	}

	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		num := generator.Intn(1000)

		select {
		case signals := <-quit:
			cancel()
			wg.Wait()
			fmt.Printf("Programm has been stopped by signal %d\n", signals)
			return
		default:
			ch <- num

		}

		time.Sleep(time.Millisecond * 100)
	}
}
