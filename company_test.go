package faker

import (
	"fmt"
	"regexp"
	"testing"
)

func TestCompanyName(t *testing.T) {
	testMatchRx(t, Company().Name, `[A-Z][a-z]+?`)
}

func TestCompanySuffix(t *testing.T) {
	testMatchRx(t, Company().Suffix, `\w+?`)
}

func TestCompanyCatchPhrase(t *testing.T) {
	testMatchRx(t, Company().CatchPhrase, `[A-z][a-z]+?`)
}

func TestCompanyBs(t *testing.T) {
	testMatchRx(t, Company().Bs, `[A-z][a-z]+?`)
}

func TestCompanyEin(t *testing.T) {
	testMatchRx(t, Company().Ein, `\d{2}-\d{7}`)
}

func TestCompanyDunsNumber(t *testing.T) {
	testMatchRx(t, Company().DunsNumber, `\d{2}-\d{3}-\d{4}`)
}

func TestCompanyLogo(t *testing.T) {
	testMatchRx(t, Company().Logo, `http://.+\d{3}.gif`)
}

func TestCompanyStringer(t *testing.T) {
	rx := `[A-Z][a-z]+?`
	res := fmt.Sprintf("%s", Company())
	if m, _ := regexp.MatchString(rx, res); !m {
		t.Errorf("expected %v to match %v", res, rx)
	}
}
