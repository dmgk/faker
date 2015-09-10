package faker

import "testing"

func TestCommerceColor(t *testing.T) {
	testMatchRx(t, Commerce().Color, `[a-z]+\.?`)
}

func TestCommerceDepartment(t *testing.T) {
	testMatchRx(t, Commerce().Department, `[A-Z][a-z]+\.?`)
}

func TestCommerceProductName(t *testing.T) {
	testMatchRx(t, Commerce().ProductName, `[A-Z][a-z]+\.?`)
}

func TestCommercePrice(t *testing.T) {
	for i := 0; i < 10; i++ {
		res := Commerce().Price()
		if res <= 0.0 || res > 100.0 {
			t.Errorf("expected %v to be in range (0.0, 100.0]", res)
		}
	}
}
