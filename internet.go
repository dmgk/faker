package faker

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

type Internet struct{}

// Example:
//	Internet{}.Email() // maritza@farrell.org
func (i Internet) Email() string {
	ss := []string{i.UserName(), i.DomainName()}
	return strings.Join(ss, "@")
}

// Example:
//	Internet{}.FreeEmail() // sven_rice@hotmail.com
func (i Internet) FreeEmail() string {
	ss := []string{i.UserName(), Fetch("internet.free_email")}
	return strings.Join(ss, "@")
}

// Example:
//	Internet{}.SafeEmail() // theron.nikolaus@example.net
func (i Internet) SafeEmail() string {
	ss := []string{
		i.UserName(),
		"example." + RandomChoice([]string{"com", "org", "net"}),
	}
	return strings.Join(ss, "@")
}

var separators []string = []string{".", "_"}
var rxNonWord = regexp.MustCompile(`\W`)

func sanitizeName(s string) string {
	return rxNonWord.ReplaceAllString(s, "")
}

// Example:
//	Internet{}.UserName() // micah_pfeffer
func (i Internet) UserName() string {
	sep := RandomChoice(separators)
	ss := []string{
		sanitizeName(Fetch("name.first_name")),
		strings.Join([]string{
			sanitizeName(Fetch("name.first_name")),
			sanitizeName(Fetch("name.last_name")),
		}, sep)}
	return strings.ToLower(RandomChoice(ss))
}

// Example:
//	Internet{}.Password(8, 16) // s5CzvVp6Ye
func (i Internet) Password(min, max int) string {
	return RandomString(RandomInt(min, max))
}

// Example:
//	Internet{}.DomainName() // rolfson.info
func (i Internet) DomainName() string {
	ss := []string{i.DomainWord(), i.DomainSuffix()}
	return strings.Join(ss, ".")
}

// Example:
//	Internet{}.DomainWord() // heller
func (i Internet) DomainWord() string {
	w := strings.Split(Fetch("company.name"), " ")[0]
	return strings.ToLower(sanitizeName(w))
}

// Example:
//	Internet{}.DomainSuffix() // net
func (i Internet) DomainSuffix() string {
	return Fetch("internet.domain_suffix")
}

// Example:
//	Internet{}.MacAddress() // 15:a9:83:29:76:26
func (i Internet) MacAddress() string {
	ss := make([]string, 6, 6)
	for i := range ss {
		ss[i] = fmt.Sprintf("%02x", rand.Int31n(256))
	}
	return strings.Join(ss, ":")
}

// Example:
//	Internet{}.IpV4Address() // 121.204.82.227
func (i Internet) IpV4Address() string {
	ss := make([]string, 4, 4)
	for i := range ss {
		ss[i] = fmt.Sprintf("%d", rand.Int31n(256))
	}
	return strings.Join(ss, ".")
}

// Example:
//	Internet{}.IpV6Address() // c697:392f:6a0e:bf6d:77e1:714a:10ab:0dbc
func (i Internet) IpV6Address() string {
	ss := make([]string, 8, 8)
	for i := range ss {
		ss[i] = fmt.Sprintf("%04x", rand.Int31n(65536))
	}
	return strings.Join(ss, ":")
}

// Example:
//	Internet{}.Url() // http://sporerhamill.net/kyla.schmitt
func (i Internet) Url() string {
	return fmt.Sprintf("http://%s/%s", i.DomainName(), i.UserName())
}

// Example:
//	Internet{}.Slug() // officiis-commodi
func (i Internet) Slug() string {
	sep := "-"
	s := strings.Join(Lorem{}.Words(2), sep)
	return rxNonWord.ReplaceAllString(s, sep)
}
