package main

import (
	"sync"
	"fmt"
	"sync/atomic"
)

var total int64

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	var i int64
	for i = 0; i < 100; i++ {
		_ = atomic.AddInt64(&total, i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()

	fmt.Println(total)
}
