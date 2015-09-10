package faker

import "testing"

func TestFinanceCreditCard(t *testing.T) {
	for _, cc := range CC_TYPES {
		for i := 0; i < 5; i++ {
			res := Finance().CreditCard(cc)
			num, checkDigit := res[:len(res)-1], res[len(res)-1:]
			luhn := luhnCheckDigit(rxDigits.ReplaceAllString(num, ""))
			if checkDigit != luhn {
				t.Errorf("expected check digit of %v to be %v", res, luhn)
			}
		}
	}
}
