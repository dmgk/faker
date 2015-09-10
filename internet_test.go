package faker

import (
	"regexp"
	"testing"
)

func TestInternetEmail(t *testing.T) {
	testMatchRx(t, Internet().Email, `\w+@\w+\.\w+`)
}

func TestInternetFreeEmail(t *testing.T) {
	testMatchRx(t, Internet().FreeEmail, `\w+@\w+\.\w+`)
}

func TestInternetSafeEmail(t *testing.T) {
	testMatchRx(t, Internet().SafeEmail, `\w+@example\.\w+`)
}

func TestInternetUserName(t *testing.T) {
	testMatchRx(t, Internet().UserName, `\w+`)
}

func TestInternetPassword(t *testing.T) {
	rx := `\w+`
	for i := 0; i < 10; i++ {
		res := Internet().Password(RandomInt(4, 8), RandomInt(8, 16))
		if m, _ := regexp.MatchString(rx, res); !m {
			t.Errorf("expected %v to match %v", res, rx)
		}
	}
}

func TestInternetDomainName(t *testing.T) {
	testMatchRx(t, Internet().DomainName, `\w+\.\w+`)
}

func TestInternetDomainWord(t *testing.T) {
	testMatchRx(t, Internet().DomainWord, `\w+`)
}

func TestInternetDomainSuffix(t *testing.T) {
	testMatchRx(t, Internet().DomainSuffix, `\w+`)
}

func TestMacAddress(t *testing.T) {
	testMatchRx(t, Internet().MacAddress, `([0-9a-f]{2}:){5}[0-9a-f]{2}`)
}

func TestIpV4Address(t *testing.T) {
	testMatchRx(t, Internet().IpV4Address, `(\d{1,3}\.){3}\d{1,3}`)
}

func TestIpV6Address(t *testing.T) {
	testMatchRx(t, Internet().IpV6Address, `([0-9a-f]{4}:){7}[0-9a-f]{4}`)
}

func TestUrl(t *testing.T) {
	testMatchRx(t, Internet().Url, `http://\w+\.\w+/\w+`)
}

func TestSlug(t *testing.T) {
	testMatchRx(t, Internet().Slug, `\w+`)
}
