package faker

type PhoneNumber struct{}

// Example:
//	PhoneNumber{}.PhoneNumber() // 1-599-267-6597 x537
func (p PhoneNumber) PhoneNumber() string {
	return Numerify(Fetch("phone_number.formats"))
}

// Example:
//	PhoneNumber{}.CellPhone() // +49-131-0003060
func (p PhoneNumber) CellPhone() string {
	return Numerify(Fetch("cell_phone.formats"))
}

// AreaCode returns random phone area code (En_US locale only).
// Example:
//	PhoneNumber{}.AreaCode() // 903
func (p PhoneNumber) AreaCode() string {
	var res string
	defer func() {
		if err := recover(); err != nil {
			res = ""
		}
	}()
	res = Fetch("phone_number.area_code")
	return res
}

// ExchangeCode returns random phone exchange code (En_US locale only).
// Example:
//	PhoneNumber{}.ExchangeCode() // 574
func (p PhoneNumber) ExchangeCode() string {
	var res string
	defer func() {
		if err := recover(); err != nil {
			res = ""
		}
	}()
	res = Fetch("phone_number.exchange_code")
	return res
}

// SubscriberNumber returns random phone subscriber number with "digits" digits.
// Example:
//	PhoneNumber{}.SubscriberNumber(4) // 1512
func (p PhoneNumber) SubscriberNumber(digits int) string {
	return Number{}.Number(digits)
}
