package cmd

//package main

import (
	"fmt"
	"module5/task_05/internal"
	"time"
)

const (
	k1   = "key1"
	step = 7
)

func main() {
	cache := internal.Cache{Storage: make(map[string]int)}

	semaphore := make(chan int, 4)

	for i := 0; i < 10; i++ {
		semaphore <- i
		go func() {
			defer func() {
				msg := <-semaphore
				fmt.Printf("first loop message - %v\n", msg)
			}()
			cache.Increase(k1, step)
			time.Sleep(time.Millisecond * 100)
		}()
	}

	for i := 0; i < 10; i++ {
		semaphore <- i
		go func(i int) {
			defer func() {
				msg := <-semaphore
				fmt.Printf("second loop message - %v\n", msg)
			}()
			cache.Set(k1, step*i)
			time.Sleep(time.Millisecond * 100)
		}(i)
	}

	for len(semaphore) > 0 {
		time.Sleep(time.Millisecond * 10)
	}
	fmt.Println(cache.Get(k1))
}
