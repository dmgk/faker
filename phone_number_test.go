package faker

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/enodata/faker/locales"
)

func TestPhoneNumberPhoneNumber(t *testing.T) {
	testMatchRx(t, PhoneNumber().PhoneNumber, `\w+`)
}

func TestPhoneNumberCellPhone(t *testing.T) {
	Locale = locales.De
	testMatchRx(t, PhoneNumber().CellPhone, `\+49-\w+`)
	Locale = locales.En
}

func TestPhoneNumberAreaCode(t *testing.T) {
	Locale = locales.En_US
	testMatchRx(t, PhoneNumber().AreaCode, `\d+`)
	Locale = locales.En
}

func TestPhoneNumberExchangeCode(t *testing.T) {
	Locale = locales.En_US
	testMatchRx(t, PhoneNumber().AreaCode, `\d+`)
	Locale = locales.En
}

func TestPhoneNumberSubscriberNumber(t *testing.T) {
	args := []int{3, 4, 5}
	for _, digits := range args {
		for i := 0; i < 10; i++ {
			rx := fmt.Sprintf(`\d{%d}`, digits)
			res := PhoneNumber().SubscriberNumber(digits)
			if m, _ := regexp.MatchString(rx, res); !m {
				t.Errorf("expected %v to match %v", res, rx)
			}
		}
	}
}

func TestPhoneNumberStringer(t *testing.T) {
	rx := `\w+`
	res := fmt.Sprintf("%s", PhoneNumber())
	if m, _ := regexp.MatchString(rx, res); !m {
		t.Errorf("expected %v to match %v", res, rx)
	}
}
