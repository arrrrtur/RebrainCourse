package internal

type Overduer struct {
	Balance int
	Debt    int
}

type Debtor interface {
	WrOffDebt() error
}

func NewCustomer(name string, age int, balance int, debt int) *Customer {
	return &Customer{
		Name: name,
		Age:  age,
		Debtor: &Overduer{
			Balance: balance,
			Debt:    debt,
		},
	}
}

func (c *Overduer) WrOffDebt() error {
	c.Debt = 0

	return nil
}

type Customer struct {
	Name string
	Age  int
	Debtor
}
