package faker

import "testing"

func TestNameName(t *testing.T) {
	testMatchRx(t, Name().Name, `(\w+\.? ?){2,3}`)
}

func TestNameFirstName(t *testing.T) {
	testMatchRx(t, Name().FirstName, `[A-Z][a-z']+`)
}

func TestNameLastName(t *testing.T) {
	testMatchRx(t, Name().LastName, `[A-Z][a-z']+`)
}

func TestNamePrefix(t *testing.T) {
	testMatchRx(t, Name().Prefix, `[A-Z][a-z]+\.?`)
}

func TestNameSuffix(t *testing.T) {
	testMatchRx(t, Name().Suffix, `[A-Z][a-z]*\.?`)
}

func TestNameTitle(t *testing.T) {
	testMatchRx(t, Name().Title, `[A-Z][a-z]+`)
}
