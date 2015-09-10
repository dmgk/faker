package faker

import "testing"

func TestHackerSaySomethingSmart(t *testing.T) {
	testMatchRx(t, Hacker().SaySomethingSmart, `[A-Z][a-z']+`)
}
