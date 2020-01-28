package factory

import (
	"fmt"
)

// PaymentMethod exported to be used globally
type PaymentMethod interface {
	pay(amount float32) string
}

type cashMethod struct {
}

func (c *cashMethod) pay(amount float32) string {
	return fmt.Sprintf("%0.2f payed using cash\n", amount)
}

type creditMethod struct {
}

func (c *creditMethod) pay(amount float32) string {
	return fmt.Sprintf("%#0.2f payed using credit card\n", amount)
}

type debitMethod struct {
}

func (c *debitMethod) pay(amount float32) string {
	return fmt.Sprintf("%#0.2f payed using debit card\n", amount)
}

// GetPaymentMethod exported to be used globally
func GetPaymentMethod(method string) (PaymentMethod, error) {
	switch method {
	case "cash":
		return new(cashMethod), nil
	case "credit":
		return new(creditMethod), nil
	case "debit":
		return new(debitMethod), nil
	default:
		return nil, fmt.Errorf("payment method %s not recognized", method)
	}
}
