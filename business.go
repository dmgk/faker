package faker

type Business struct{}

// Example:
//  Business{}.CreditCardNumber() // 1234-2121-1221-1211
func (b Business) CreditCardNumber() string {
	return Fetch("business.credit_card_numbers")
}

// Example:
//  Business{}.CreditCardExpiryDate() // 2015-11-11
func (b Business) CreditCardExpiryDate() string {
	return Fetch("business.credit_card_expiry_dates")
}

// Example:
//  Business{}.CreditCardType() // mastercard
func (b Business) CreditCardType() string {
	return Fetch("business.credit_card_types")
}
