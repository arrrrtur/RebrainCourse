package cmd

//package main

import (
	"fmt"
	"module4/03_task/internal"
)

func main() {
	customer := internal.NewCustomer("Dmitry", "23", 10_000, 0, true)
	fmt.Println(internal.CalcPrice(customer, 800))
}
