package main

import (
	"fmt"
	"sync"
	"time"
)

type sharedMap struct {
	vault map[string]any
	mu    *sync.Mutex
}

func (sm *sharedMap) set(key string, value any) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.vault[key] = value
}

func main() {
	// using mutex
	var mu sync.Mutex

	map1 := sharedMap{
		vault: make(map[string]any, 10),
		mu:    &mu,
	}

	go map1.set("key1", "value1")
	go map1.set("key2", "value2")
	go map1.set("key3", 123)

	time.Sleep(100 * time.Millisecond)

	fmt.Println(map1.vault)

	// using sync.Map
	var map2 sync.Map
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			map2.Store(n, n*10)
		}(i)
	}

	wg.Wait()

	for i := 0; i < 3; i++ {
		v, _ := map2.Load(i)
		fmt.Printf("Key: %d, value: %d\n", i, v)
	}
}
