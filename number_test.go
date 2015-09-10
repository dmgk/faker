package faker

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

func TestNumberNumber(t *testing.T) {
	digits := []int{1, 5, 17, 37}
	for _, d := range digits {
		rx := fmt.Sprintf(`\d{%d}`, d)
		for i := 0; i < 10; i++ {
			res := Number().Number(d)
			if m, _ := regexp.MatchString(rx, res); !m {
				t.Errorf("expected %v to match %v", res, rx)
			}
		}
	}
}

func TestNumberDecimal(t *testing.T) {
	args := []struct {
		precision, scale int
	}{
		{3, 1},
		{3, 3},
		{8, 2},
		{8, 7},
		{10, 0},
	}
	for _, a := range args {
		rx := fmt.Sprintf(`\d{%d}\.\d{%d}`, a.precision-a.scale, a.scale)
		for i := 0; i < 10; i++ {
			res := Number().Decimal(a.precision, a.scale)
			if m, _ := regexp.MatchString(rx, res); !m {
				t.Errorf("expected %v to match %v", res, rx)
			}
		}
	}
}

func TestNumberDigit(t *testing.T) {
	testMatchRx(t, Number().Digit, `\d`)
}

func TestNumberHexadecimal(t *testing.T) {
	digits := []int{1, 7, 19, 4, 8}
	for _, d := range digits {
		rx := fmt.Sprintf(`[0-9a-f]{%d}`, d)
		for i := 0; i < 10; i++ {
			res := Number().Hexadecimal(d)
			if m, _ := regexp.MatchString(rx, res); !m {
				t.Errorf("expected %v to match %v", res, rx)
			}
		}
	}
}

func TestNumberBetween(t *testing.T) {
	args := []struct {
		min, max int
	}{
		{-100, 100},
		{-5000, -5},
		{0, 5000},
		{10, 10},
		{0, 0},
	}
	rx := `-?\d+`
	for _, a := range args {
		for i := 0; i < 10; i++ {
			res := Number().Between(a.min, a.max)
			ires, _ := strconv.Atoi(res)
			if ires < a.min || ires > a.max {
				t.Errorf("expected %v to be in [%d, %d]", res, a.min, a.max)
			}
			if m, _ := regexp.MatchString(rx, res); !m {
				t.Errorf("expected %v to match %v", res, rx)
			}
		}
	}
}

func TestNumberPositive(t *testing.T) {
	maxs := []int{10, 1000000, 0, 1}
	rx := `\d+`
	for _, max := range maxs {
		for i := 0; i < 10; i++ {
			res := Number().Positive(max)
			ires, _ := strconv.Atoi(res)
			if ires < 0 || ires > max {
				t.Errorf("expected %v to be in [0, %d]", res, max)
			}
			if m, _ := regexp.MatchString(rx, res); !m {
				t.Errorf("expected %v to match %v", res, rx)
			}
		}
	}
}

func TestNumberNegative(t *testing.T) {
	mins := []int{-42, 0, -10000, -1}
	rx := `\d+`
	for _, min := range mins {
		for i := 0; i < 10; i++ {
			res := Number().Negative(min)
			ires, _ := strconv.Atoi(res)
			if ires < min || ires > 0 {
				t.Errorf("expected %v to be in [%d, 0]", res, min)
			}
			if m, _ := regexp.MatchString(rx, res); !m {
				t.Errorf("expected %v to match %v", res, rx)
			}
		}
	}
}
