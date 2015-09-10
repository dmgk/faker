package faker

import "testing"

func TestBusinessCreditCardNumber(t *testing.T) {
	list := valueAt("business.credit_card_numbers").([]string)
	for i := 0; i < 5; i++ {
		res := Business().CreditCardNumber()
		if !includesString(list, res) {
			t.Errorf("expected %v to include %v", list, res)
		}
	}
}

func TestBusinessCreditCardExpiryDate(t *testing.T) {
	list := valueAt("business.credit_card_expiry_dates")
	for i := 0; i < 5; i++ {
		res := Business().CreditCardExpiryDate()
		if !includesString(list.([]string), res) {
			t.Errorf("expected %v to include %v", list, res)
		}
	}
}

func TestBusinessCreditCardType(t *testing.T) {
	list := valueAt("business.credit_card_types")
	for i := 0; i < 5; i++ {
		res := Business().CreditCardType()
		if !includesString(list.([]string), res) {
			t.Errorf("expected %v to include %v", list, res)
		}
	}
}
