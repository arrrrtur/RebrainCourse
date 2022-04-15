package _2_task

//package main

import (
	"fmt"
	"math"
)

func main() {
	// task2
	A := new(int)
	B := 123
	A = &B

	fmt.Printf("task2:\n\t%d\n", *A)

	*A = 321

	fmt.Printf("\n\t%d\n", B)

	// task3
	R := new(float64)
	*R = 35 / (2 * math.Pi)
	square := math.Round((math.Pi*math.Pow(*R, 2))*100) / 100
	fmt.Printf("task3:\n\t%.2f\n", square)

}
