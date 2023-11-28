package main

import (
	"fmt"
	"time"
)

const (
	limit = 2
)

func main() {
	waitQueue := make(chan int, limit)
	go func() {
		var c int
		fn := func(c *int, n int) {
			defer func() { *c-- }()
			fmt.Println("counter:", n)
			time.Sleep(1 * time.Second)
		}
		for n := range waitQueue {
			for c >= limit {
			}
			go fn(&c, n)
			c++
		}
	}()

	for i := 1; i <= 10; i++ {
		waitQueue <- i
	}

	time.Sleep(10 * time.Second)
}
