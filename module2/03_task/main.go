package _3_task

//package main

import "fmt"

func main() {
	// task2
	week := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	workingWeek := week[:5]
	week = week[5:]

	fmt.Printf("task2: \n\tWorking days - %v,\n\tHolidays - %v\n", workingWeek, week)

	// task3
	fullWeek := append(workingWeek, week...)
	fmt.Printf("Task3: \n\tFull week - %v\n", fullWeek)
}
