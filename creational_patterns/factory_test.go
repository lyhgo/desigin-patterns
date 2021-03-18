package creational_patterns

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A payment metheod of type 'Cash' must exist")
	}
	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestCreatePaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatal("A payment metheod of type 'DebitCard' must exist")
	}
	msg := payment.Pay(20.30)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The debitCard payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestCreatePaymentMethodNonExistent(t *testing.T) {
	_, err := GetPaymentMethod(20)
	if err == nil {
		t.Error("A payment metheod of ID '20' must exist")
	}
	t.Log("LOG:", err)
}
