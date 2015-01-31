package faker

type Name struct{}

// Example:
//	Name{}.Name() // Natasha Hartmann
func (n Name) Name() string {
	return Fetch("name.name")
}

// Example:
//	Name{}.FirstName() // Carolina
func (n Name) FirstName() string {
	return Fetch("name.first_name")
}

// Example:
//	Name{}.LastName() // Kohler
func (n Name) LastName() string {
	return Fetch("name.last_name")
}

// Example:
//	Name{}.Prefix() // Dr.
func (n Name) Prefix() string {
	return Fetch("name.prefix")
}

// Example:
//	Name{}.Suffix() // Jr.
func (n Name) Suffix() string {
	return Fetch("name.suffix")
}

// Example:
//	Name{}.Title() // Chief Functionality Orchestrator
func (n Name) Title() string {
	return Fetch("name.title.descriptor") + " " + Fetch("name.title.level") + " " + Fetch("name.title.job")
}

// String returns a random full name.
// Example:
//	fmt.Println(Name{}) //
func (n Name) String() string {
	return n.Name()
}
