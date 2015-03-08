package faker

import (
	"fmt"
	"regexp"
	"testing"
)

func TestNumerify(t *testing.T) {
	examples := []struct {
		s  string
		rx string
	}{
		{"#", `\A\d\z`},
		{"###", `\A\d{3}\z`},
		{"##-###", `\A\d{2}-\d{3}\z`},
		{"catch##", `\Acatch\d{2}\z`},
	}

	for _, x := range examples {
		res := Numerify(x.s)
		if m, _ := regexp.MatchString(x.rx, res); !m {
			t.Errorf("expected %q (%q) to match %v", x.s, res, x.rx)
		}
	}
}

func TestLetterify(t *testing.T) {
	examples := []struct {
		s  string
		rx string
	}{
		{"?", `\A[A-Z]\z`},
		{"???", `\A[A-Z]{3}\z`},
		{"??0???", `\A[A-Z]{2}0[A-Z]{3}\z`},
		{"42???", `\A42[A-Z]{3}\z`},
	}

	for _, x := range examples {
		res := Letterify(x.s)
		if m, _ := regexp.MatchString(x.rx, res); !m {
			t.Errorf("expected %q (%q) to match %v", x.s, res, x.rx)
		}
	}
}

func TestRegexify(t *testing.T) {
	// US phone number
	rx := `^(1-?)[2-8][0-1][0-9]-\d{3}-\d{4}$`
	for i := 0; i < 30; i++ {
		res, err := Regexify("/" + rx + "/")
		if err != nil {
			t.Fatal(err)
		}

		if m, _ := regexp.MatchString(rx, res); !m {
			t.Errorf("expected %q to match %v", res, rx)
		}
	}

	// UK post code
	rx = `^([A-PR-UWYZ0-9][A-HK-Y0-9][AEHMNPRTVXY0-9]?[ABEHMNPRVWXY0-9]? {1,2}[0-9][ABD-HJLN-UW-Z]{2}|GIR 0AA)$`
	for i := 0; i < 30; i++ {
		res, err := Regexify(rx)
		if err != nil {
			t.Fatal(err)
		}

		if m, _ := regexp.MatchString(rx, res); !m {
			t.Errorf("expected %q to match %v", res, rx)
		}
	}
}

func TestFetch(t *testing.T) {
	rx := `\w+`
	for i := 0; i < 10; i++ {
		res := Fetch("company.bs")
		if m, _ := regexp.MatchString(rx, res); !m {
			t.Errorf("expected %v to match %v", res, rx)
		}
	}
}

func TestFetchInterpolation(t *testing.T) {
	rx := `\w+[\s-,]+\w+`
	for i := 0; i < 10; i++ {
		res := Fetch("company.name")
		if m, _ := regexp.MatchString(rx, res); !m {
			t.Errorf("expected %v to match %v", res, rx)
		}
	}
}

func TestFetchInvalidKeyPathPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			// ignore
		}
	}()
	Fetch("non.existing.key.path")
	t.Error("expected invalid key path panic")
}

func TestFetchInvalidNodeTypePanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			// ignore
		}
	}()
	Fetch("name")
	t.Error("expected invalid value type panic")
}

// Testing helpers

func testMatchRxN(t *testing.T, f func() string, rx string, n int, println bool) {
	for i := 0; i < 1000; i++ {
		res := f()
		if m, _ := regexp.MatchString(rx, res); !m {
			t.Errorf("expected %v to match %v", res, rx)
		}
		if println {
			fmt.Println(res)
		}
	}
}

func testMatchRx(t *testing.T, f func() string, rx string) {
	testMatchRxN(t, f, rx, 10, false)
}

func testMatchRxP(t *testing.T, f func() string, rx string) {
	testMatchRxN(t, f, rx, 10, true)
}
