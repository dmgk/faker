package faker

import (
	crand "crypto/rand"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
)

type Number struct{}

// Number returns random decimal number "digits" digits long.
// Example:
//	Number{}.Number(5) // 43202
func (n Number) Number(digits int) string {
	if digits <= 0 {
		panic("invalid digits value")
	}
	dd := make([]string, digits, digits)
	for i := range dd {
		dd[i] = n.Digit()
	}
	return strings.Join(dd, "")
}

// Decimal returns random decimal number with "precision" precision and "scale" scale.
// Example:
//	Number{}.Decimal(8, 2) // 879420.60
func (n Number) Decimal(precision, scale int) string {
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

// Digit returns random decimal digit.
// Example:
//	Number{}.Digit() // 7
func (n Number) Digit() string {
	return fmt.Sprintf("%d", rand.Int31n(10))
}

// Hexadecimal returns random hex number "digits" hexdigits long.
// Example:
//	Number{}.Hexadecimal(4) // e7f3
func (n Number) Hexadecimal(digits int) string {
	if digits <= 0 {
		panic("invalid digits value")
	}
	bytes := make([]byte, (digits+1)/2)
	crand.Read(bytes)
	return hex.EncodeToString(bytes)[:digits]
}

// Between returns random decimal number in [min, max] range.
// Example:
//	Number{}.Between(-100, 100) // -47
func (n Number) Between(min, max int) string {
	if min > max {
		panic("invalid range")
	}
	return fmt.Sprintf("%d", RandomInt(min, max))
}

// Positive returns random decimal number in [0, max] range.
// Example:
//	Number{}.Positive(10) // 3
func (n Number) Positive(max int) string {
	if max < 0 {
		panic("invalid max value")
	}
	return n.Between(0, max)
}

// Positive returns random decimal number in [min, 0] range.
// Example:
//	Number{}.Negative(-42) // -16
func (n Number) Negative(min int) string {
	if min > 0 {
		panic("invalid max value")
	}
	return n.Between(min, 0)
}
