package internal

type Customer struct {
	Name         string
	Age          string
	Balance      int
	Debt         int
	Discount     bool
	CalcDiscount func() (int, error)
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

func CalcPrice(customer Customer, price int) (int, error) {
	discount, err := customer.CalcDiscount()
	if err != nil {
		return 0, err
	}
	return price - discount, nil
}
