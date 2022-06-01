package cmd

//package main

import (
	"fmt"
	"module5/05_task/internal"
	"time"
)

const (
	k1   = "key1"
	step = 7
)

func main() {
	cache := internal.Cache{Storage: make(map[string]int)}

	semaphore := internal.NewSemaphore(4)

	for i := 0; i < 10; i++ {
		if err := semaphore.Acquire(); err != nil {
			panic(fmt.Errorf("err: %+v\n", err.Error()))
		}
		go func() {
			defer func() {
				if err := semaphore.Release(); err != nil {
					panic(fmt.Errorf("err: %+v\n", err.Error()))
				}
			}()
			cache.Increase(k1, step)
			time.Sleep(time.Millisecond * 100)
		}()
	}

	for i := 0; i < 10; i++ {
		if err := semaphore.Acquire(); err != nil {
			panic(fmt.Errorf("err: %+v\n", err.Error()))

		}
		go func(i int) {
			defer func() {
				if err := semaphore.Release(); err != nil {
					panic(fmt.Errorf("err: %+v\n", err.Error()))
				}
			}()
			cache.Set(k1, step*i)
			time.Sleep(time.Millisecond * 100)
		}(i)
	}

	for semaphore.Len() > 0 {
		time.Sleep(time.Millisecond * 10)
	}
	fmt.Println(cache.Get(k1))
}
