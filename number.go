package faker

import (
	crand "crypto/rand"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
)

type FakeNumber interface {
	Number(digits int) string            // => "43202"
	Decimal(precision, scale int) string // => "879420.60"
	Digit() string                       // => "7"
	Hexadecimal(digits int) string       // => "e7f3"
	Between(min, max int) string         // => "-47"
	Positive(max int) string             // => "3"
	Negative(min int) string             // => "-16"
}

type fakeNumber struct{}

func Number() FakeNumber {
	return fakeNumber{}
}

func (n fakeNumber) Number(digits int) string {
	if digits <= 0 {
		panic("invalid digits value")
	}
	dd := make([]string, digits, digits)
	for i := range dd {
		dd[i] = n.Digit()
	}
	return strings.Join(dd, "")
}

func (n fakeNumber) Decimal(precision, scale int) string {
	if precision <= 0 || scale < 0 || precision < scale {
		panic("invalid precision or scale values")
	}
	s := n.Number(precision)
	ii := s[:precision-scale]
	if len(ii) == 0 {
		ii = "0"
	}
	ff := s[precision-scale:]
	if len(ff) == 0 {
		ff = "0"
	}

	return ii + "." + ff
}

func (n fakeNumber) Digit() string {
	return fmt.Sprintf("%d", rand.Int31n(10))
}

func (n fakeNumber) Hexadecimal(digits int) string {
	if digits <= 0 {
		panic("invalid digits value")
	}
	bytes := make([]byte, (digits+1)/2)
	crand.Read(bytes)
	return hex.EncodeToString(bytes)[:digits]
}

func (n fakeNumber) Between(min, max int) string {
	if min > max {
		panic("invalid range")
	}
	return fmt.Sprintf("%d", RandomInt(min, max))
}

func (n fakeNumber) Positive(max int) string {
	if max < 0 {
		panic("invalid max value")
	}
	return n.Between(0, max)
}

func (n fakeNumber) Negative(min int) string {
	if min > 0 {
		panic("invalid max value")
	}
	return n.Between(min, 0)
}
