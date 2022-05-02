package cmd

import (
	"fmt"
	"module4/02_task/internal"
)

func main() {
	customer := internal.NewCustomer("Dmitry", "23", 10_000, 0, true)
	fmt.Println(internal.CalcPrice(*customer, 800))
}
