package cmd

//package main

import (
	"fmt"
	"module5/04_task/internal"
	"sync"
	"time"
)

const (
	k1   = "key1"
	step = 7
)

func main() {
	cache := internal.Cache{Storage: make(map[string]int)}

	var waitGroup sync.WaitGroup
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			cache.Increase(k1, step)
			time.Sleep(time.Millisecond * 100)
		}()
	}

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			cache.Set(k1, step*i)
			time.Sleep(time.Millisecond * 100)
		}(i)
	}
	waitGroup.Wait()
	fmt.Println(cache.Get(k1))
}
