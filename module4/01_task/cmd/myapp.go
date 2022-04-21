package cmd

import (
	"errors"
	"fmt"
	"module04/01_task/internal"
)

const DEFAULT_DISCOUNT = 500

func main() {
	customer := internal.NewCustomer("Dmitry", "23", 10_000, 1000, true)
	customer.CalcDiscount = func() (int, error) {
		if !customer.Discount {
			return 0, errors.New("discount not available")
		}
		result := DEFAULT_DISCOUNT - customer.Debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	}

	fmt.Println(internal.CalcPrice(*customer, 800))

}
