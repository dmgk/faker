package faker

import "testing"

func TestTeamName(t *testing.T) {
	testMatchRx(t, Team{}.Name, `[A-Z][a-z]+`)
}

func TestTeamCreature(t *testing.T) {
	testMatchRx(t, Team{}.Creature, `\w+`)
}

func TestTeamState(t *testing.T) {
	testMatchRx(t, Team{}.State, `[A-Z][a-z]+`)
}
