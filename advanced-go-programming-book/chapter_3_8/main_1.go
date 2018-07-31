package main

import (
	"github.com/petermattis/goid"
	"time"
	"fmt"
)

func main()  {
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000000; j++ {
				fmt.Printf("[#%d] %d", goid.Get(), j)
				time.Sleep(10e9)
			}
		}()
	}
	select {}
}