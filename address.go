package faker

import (
	"fmt"
	"math/rand"
	"reflect"
)

type Address struct{}

// Example:
//  Address{}.City() // North Dessie
func (a Address) City() string {
	return Fetch("address.city")
}

// Example:
//  Address{}.StreetName() // Buckridge Lakes
func (a Address) StreetName() string {
	return Fetch("address.street_name")
}

// Example:
//  Address{}.StreetAddress() // 586 Sylvester Turnpike
func (a Address) StreetAddress() string {
	return Numerify(Fetch("address.street_address"))
}

// Example:
//  Address{}.SecondaryAddress() // Apt. 411
func (a Address) SecondaryAddress() string {
	return Numerify(Fetch("address.secondary_address"))
}

// Example:
//  Address{}.BuildingNumber() // 754
func (a Address) BuildingNumber() string {
	return NumerifyAndLetterify(Fetch("address.building_number"))
}

// Example:
//  Address{}.Postcode() // 31340
func (a Address) Postcode() string {
	return NumerifyAndLetterify(Fetch("address.postcode"))
}

// Example:
//  Address{}.PostcodeByState("AZ") // 46511
func (a Address) PostcodeByState(state string) string {
	// postcode_by_state can be either a map[string] or a slice (as in default En locale)
	switch pbs := valueAt("address.postcode_by_state").(type) {
	case map[string]interface{}:
		_, ok := pbs[state]
		if ok {
			return NumerifyAndLetterify(Fetch("address.postcode_by_state." + state))
		} else {
			panic(fmt.Sprintf("invalid state: %v", state))
		}
	case []string:
		return NumerifyAndLetterify(Fetch("address.postcode_by_state"))
	default:
		panic(fmt.Sprintf("invalid postcode_by_state value type: %v", reflect.TypeOf(pbs)))
	}
}

// ZipCode is an alias for Postcode.
func (a Address) ZipCode() string {
	return a.Postcode()
}

// ZipCodeByState is an alias for PostcodeByState
func (a Address) ZipCodeByState(state string) string {
	return a.PostcodeByState(state)
}

// Example:
//  Address{}.TimeZone() // Asia/Taipei
func (a Address) TimeZone() string {
	return Fetch("address.time_zone")
}

// Example:
//  Address{}.CityPrefix() // East
func (a Address) CityPrefix() string {
	return Fetch("address.city_prefix")
}

// Example:
//  Address{}.CitySuffix() // town
func (a Address) CitySuffix() string {
	return Fetch("address.city_suffix")
}

// Example:
//  Address{}.StreetSuffix() // Square
func (a Address) StreetSuffix() string {
	return Fetch("address.street_suffix")
}

// Example:
//  Address{}.State() // Maryland
func (a Address) State() string {
	return Fetch("address.state")
}

// Example:
//  Address{}.StateAbbr() // IL
func (a Address) StateAbbr() string {
	return Fetch("address.state_abbr")
}

// Example:
//  Address{}.Country() // Uruguay
func (a Address) Country() string {
	return Fetch("address.country")
}

// Example:
//  Address{}.CountryCode() // JP
func (a Address) CountryCode() string {
	return Fetch("address.country_code")
}

// Example:
//  Address{}.Latitude() // -38.811367
func (a Address) Latitude() float32 {
	return (rand.Float32() * 180.0) - 90.0
}

// Example:
//  Address{}.Longitude() // 89.2171
func (a Address) Longitude() float32 {
	return (rand.Float32() * 360.0) - 180.0
}

// String returns US style formatted address.
// Example:
//  fmt.Println(Address{}) // 6071 Heaney Island Suite 553, Ebbaville Texas 37307
func (a Address) String() string {
	return fmt.Sprintf("%v %v, %v %v %v", a.StreetAddress(), a.SecondaryAddress(), a.City(), a.State(), a.Postcode())
}
