package internal

import "errors"

const DEFAULT_DISCOUNT = 500

type Discounter interface {
	CalcDiscount() (int, error)
}

type Customer struct {
	Name     string
	Age      string
	Balance  int
	Debt     int
	Discount bool
}

func (customer *Customer) CalcDiscount() (int, error) {
	if !customer.Discount {
		return 0, errors.New("discount not available")
	}
	result := DEFAULT_DISCOUNT - customer.Debt
	if result < 0 {
		return 0, nil
	}
	return result, nil
}

func NewCustomer(Name, Age string, Balance, Debt int, Discount bool) *Customer {
	return &Customer{
		Name:     Name,
		Age:      Age,
		Balance:  Balance,
		Debt:     Debt,
		Discount: Discount,
	}
}

func CalcPrice(discounter Discounter, price int) (int, error) {
	discount, err := discounter.CalcDiscount()
	if err != nil {
		return 0, err
	}
	return price - discount, nil
}
