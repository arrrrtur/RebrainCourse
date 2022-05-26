package cmd

//package main

import (
	"fmt"
	"module4/05_task/internal"
)

func main() {
	partner := internal.NewCustomer("Dmitry", 23, 10000, 1000)
	startTransaction(partner)
	fmt.Printf("%+v\n", partner)
}

func startTransaction(debtor internal.Debtor) error {
	return debtor.WrOffDebt()
}
