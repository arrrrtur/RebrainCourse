package cmd

//package main

import (
	"fmt"
	"module5/01_task/internal"
	"time"
)

func main() {
	go internal.Spinner(100 * time.Millisecond)
	go fmt.Printf("fibonachi of 44 is %v\n", internal.Fib(44))
	go fmt.Printf("fibonachi of 45 is %v\n", internal.Fib(45))
	time.Sleep(time.Millisecond)
}
