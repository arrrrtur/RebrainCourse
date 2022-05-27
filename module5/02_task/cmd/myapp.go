package cmd

//package main

import (
	"fmt"
	"module5/02_task/internal"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1) // max processors
	go internal.Spinner(100 * time.Millisecond)
	go fmt.Printf("fibonachi of 44 is %v\n", internal.Fib(44))
	go fmt.Printf("fibonachi of 45 is %v\n", internal.Fib(45))
	time.Sleep(time.Millisecond)
}
