package _1_task

//package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	num1 := "104"
	num2 := 35

	stringToNum, _ := strconv.Atoi(num1)
	numToString := strconv.Itoa(num2)

	fmt.Printf("string to number - %d,\nnumber to string - %s\n", stringToNum, numToString)

	type AmericanVelocity float64
	type EuropeanVelocity float64

	// 120.4 m\s to km\h
	result1 := EuropeanVelocity(math.Round((120.4*60*60/1000)*100) / 100)
	// 130 m\s to mile\h
	result2 := AmericanVelocity(math.Round((130*60*60/(1.609*1000))*100) / 100)

	fmt.Printf("task 7: %.2f - %.2f", result1, result2)

}
