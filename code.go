package faker

import (
	"fmt"
	"strconv"
)

type Code struct{}

// Example:
//  Code{}.Isbn10() // 048931033-8
func (c Code) Isbn10() string {
	val, err := Regexify(`\d{9}`)
	if err != nil {
		panic(err)
	}

	var sum int = 0
	for i, v := range val {
		n, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err)
		}
		sum += (i + 1) * n
	}
	rem := sum % 11

	if rem == 10 {
		return fmt.Sprintf("%s-X", val)
	} else {
		return fmt.Sprintf("%s-%d", val, rem)
	}
}

func ean13() (ean string, checkDigit string) {
	ean, err := Regexify(`\d{12}`)
	if err != nil {
		panic(err)
	}

	var sum int = 0
	for i, d := range ean {
		n, err := strconv.Atoi(string(d))
		if err != nil {
			panic(err)
		}
		if i%2 == 0 {
			sum += n
		} else {
			sum += n * 3
		}
	}
	rem := sum % 10
	checkDigit = strconv.Itoa((10 - rem) % 10)
	return
}

// Example:
//  Code{}.Isbn13() // 391668236072-1
func (c Code) Isbn13() string {
	ean, checkDigit := ean13()
	return fmt.Sprintf("%s-%s", ean, checkDigit)
}

// Example:
//  Code{}.Ean13() // 7742864258656
func (c Code) Ean13() string {
	ean, checkDigit := ean13()
	return fmt.Sprintf("%s%s", ean, checkDigit)
}

// Example:
//  Code{}.Ean8() // 03079010
func (c Code) Ean8() string {
	ean, err := Regexify(`\d{7}`)
	if err != nil {
		panic(err)
	}

	var sum int = 0
	for i, d := range ean {
		n, err := strconv.Atoi(string(d))
		if err != nil {
			panic(err)
		}
		if i%2 == 0 {
			sum += n * 3
		} else {
			sum += n
		}
	}
	rem := sum % 10

	return fmt.Sprintf("%s%d", ean, (10-rem)%10)
}

var rutFactors = [...]int{3, 2, 7, 6, 5, 4, 3, 2}

// Example:
//  Code{}.Rut() // 14371602-3
func (c Code) Rut() string {
	rut, err := Regexify(`\d{8}`)
	if err != nil {
		panic(err)
	}

	var sum int = 0
	for i, d := range rut {
		n, err := strconv.Atoi(string(d))
		if err != nil {
			panic(err)
		}
		sum += n * rutFactors[i]
	}
	rem := 11 - (sum % 11)

	var checkDigit string
	switch rem {
	case 11:
		checkDigit = "0"
	case 10:
		checkDigit = "K"
	default:
		checkDigit = strconv.Itoa(rem)
	}

	return fmt.Sprintf("%s-%s", rut, checkDigit)
}
