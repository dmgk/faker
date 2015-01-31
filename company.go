package faker

import "fmt"

type Company struct{}

// Example:
//  Company{}.Name() // Aufderhar LLC
func (c Company) Name() string {
	return Fetch("company.name")
}

// Example:
//  Company{}.Suffix() // Inc
func (c Company) Suffix() string {
	return Fetch("company.suffix")
}

// Example:
//  Company{}.CatchPhrase() // Universal logistical artificial intelligence
func (c Company) CatchPhrase() string {
	return Fetch("company.buzzwords")
}

// Example:
//  Company{}.Bs() // engage distributed applications
func (c Company) Bs() string {
	return Fetch("company.bs")
}

// Example:
//  Company{}.Ein() // 58-6520513
func (c Company) Ein() string {
	return Numerify("##-#######")
}

// Example:
//  Company{}.DunsNumber() // 16-708-2968
func (c Company) DunsNumber() string {
	return Numerify("##-###-####")
}

// Example:
//  Company{}.Logo() // http://www.biz-logo.com/examples/015.gif
func (c Company) Logo() string {
	n := RandomInt(1, 77)
	return fmt.Sprintf("http://www.biz-logo.com/examples/%03d.gif", n)
}

// Example:
//  fmt.Println(Company{}) // Jaskolski-Morar
func (c Company) String() string {
	return c.Name()
}
