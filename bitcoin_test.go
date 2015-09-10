package faker

import (
	"fmt"
	"regexp"
	"testing"
)

func TestBitcoinAddress(t *testing.T) {
	testMatchRx(t, Bitcoin().Address, `^[13][1-9A-Za-z][^OIl]{20,40}`)
}

func TestBitcoinStringer(t *testing.T) {
	rx := `^[13][1-9A-Za-z][^OIl]{20,40}`
	res := fmt.Sprintf("%s", Bitcoin())
	if m, _ := regexp.MatchString(rx, res); !m {
		t.Errorf("expected %v to match %v", res, rx)
	}
}
