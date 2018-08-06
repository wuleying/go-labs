package main

import (
	"fmt"
)

func main() {
	num := make([]int, 5)
	for i := 0; i < len(num); i++ {
		num[i] = i * i
	}
	fmt.Println(num)
}
