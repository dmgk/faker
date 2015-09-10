package faker

import (
	"fmt"
	"regexp"
	"testing"
)

func TestAddressCity(t *testing.T) {
	testMatchRx(t, Address().City, `[A-Z][a-z']+`)
}

func TestAddressStreetName(t *testing.T) {
	testMatchRx(t, Address().StreetName, `[A-Z][a-z']+`)
}

func TestAddressStreetAddress(t *testing.T) {
	testMatchRx(t, Address().StreetAddress, `\d+\s[A-Z][a-z']+`)
}

func TestAddressSecondaryAddress(t *testing.T) {
	testMatchRx(t, Address().SecondaryAddress, `[A-Z][a-z]+\.?\s\d+`)
}

func TestAddressBuildingNumber(t *testing.T) {
	testMatchRx(t, Address().BuildingNumber, `\d+`)
}

func TestAddressPostcode(t *testing.T) {
	testMatchRx(t, Address().Postcode, `[\d-]+`)
}

func TestAddressPostcodeByState(t *testing.T) {
	rx := `[\d]+(-\d+)?`
	for i := 0; i < 10; i++ {
		res := Address().PostcodeByState("AZ")
		if m, _ := regexp.MatchString(rx, res); !m {
			t.Errorf("expected %v to match %v", res, rx)
		}
	}
}

func TestAddressTimeZone(t *testing.T) {
	testMatchRx(t, Address().TimeZone, `[A-Z]\w+/[A-Z]\w+`)
}

func TestAddressCityPrefix(t *testing.T) {
	testMatchRx(t, Address().CityPrefix, `\w+`)
}

func TestAddressCitySuffix(t *testing.T) {
	testMatchRx(t, Address().CitySuffix, `\w+`)
}

func TestAddressStreetSuffix(t *testing.T) {
	testMatchRx(t, Address().StreetSuffix, `\w+`)
}

func TestAddressState(t *testing.T) {
	testMatchRx(t, Address().State, `\w+`)
}

func TestAddressStateAbbr(t *testing.T) {
	testMatchRx(t, Address().StateAbbr, `\w+`)
}

func TestCountry(t *testing.T) {
	testMatchRx(t, Address().Country, `\w+`)
}

func TestCountryCode(t *testing.T) {
	testMatchRx(t, Address().CountryCode, `\w+`)
}

func TestLatitude(t *testing.T) {
	for i := 0; i < 10; i++ {
		res := Address().Latitude()
		if res < -90.0 || res > 90.0 {
			t.Errorf("expected %v to be between -90.0 and 90.0", res)
		}
	}
}

func TestLongitude(t *testing.T) {
	for i := 0; i < 10; i++ {
		res := Address().Longitude()
		if res < -180.0 || res > 180.0 {
			t.Errorf("expected %v to be between -180.0 and 180.0", res)
		}
	}
}

func TestAddressStringer(t *testing.T) {
	rx := `[A-Z][a-z]+`
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("%v", Address())
		if m, _ := regexp.MatchString(rx, res); !m {
			t.Errorf("expected %v to match %v", res, rx)
		}
	}
}
