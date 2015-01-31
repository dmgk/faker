package faker

type Team struct{}

// Example:
//  Team{}.Name() // Colorado cats
func (t Team) Name() string {
	return Fetch("team.name")
}

// Example:
//  Team{}.Creature() // cats
func (t Team) Creature() string {
	return Fetch("team.creature")
}

// Example:
//  Team{}.State() // Oregon
func (t Team) State() string {
	return Fetch("address.state")
}
