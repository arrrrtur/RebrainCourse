//package main
package cmd

import (
	"fmt"
	"module4/04_task/internal"
)

func main() {
	partner := internal.NewPartner("Dmitry", 23, 10000, 1000)
	startTransaction(partner)
	fmt.Printf("%+v\n", partner)
}

func startTransaction(debtor internal.Debtor) error {
	return debtor.WrOffDebt()
}

func startTransactionDynamic(debtor internal.Debtor) error {
	return debtor.WrOffDebt()
}
