package _5_task

//package main

import "fmt"

func main() {
	// task3
	sl := []string{"bebebe", "bababa", "bububu"}
	fmt.Printf("task3:\n\tsl contains bububu - %b\n", contains(sl, "bububu"))

	//task4
	nums := []int{1, 4, 6, 2, 7, 34, 13}
	fmt.Printf("task4:\n\tmax in %v is %d\n", nums, getMax(nums...))
}

func contains(a []string, x string) bool {
	for _, str := range a {
		if x == str {
			return true
		}
	}
	return false
}

func getMax(sl ...int) int {
	var max int
	for _, num := range sl {
		if max == 0 {
			max = num
		}
		if max < num {
			max = num
		}
	}

	return max
}
