package factory

import "testing"

import "strings"

func TestCashMethod(t *testing.T) {
	payment, err := GetPaymentMethod("cash")
	if err != nil {
		t.Fatal(err)
	}
	msg := payment.pay(10.20)
	if !strings.Contains(msg, "payed using cash") {
		t.Error("The cash payment method message wasn't correct")
	}
	t.Log("LOG::", msg)
}

func TestCreditMethod(t *testing.T) {
	payment, err := GetPaymentMethod("credit")
	if err != nil {
		t.Fatal(err)
	}
	msg := payment.pay(40.20)
	if !strings.Contains(msg, "payed using credit card") {
		t.Error("The credit card payment method message wasn't correct")
	}
	t.Log("LOG::", msg)
}

func TestDebitMethod(t *testing.T) {
	payment, err := GetPaymentMethod("debit")
	if err != nil {
		t.Fatal(err)
	}
	msg := payment.pay(50.20)
	if !strings.Contains(msg, "payed using debit card") {
		t.Error("The debit card payment method message wasn't correct")
	}
	t.Log("LOG::", msg)

}
