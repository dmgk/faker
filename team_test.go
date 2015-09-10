package faker

import (
	"fmt"
	"regexp"
	"testing"
)

func TestTeamName(t *testing.T) {
	testMatchRx(t, Team().Name, `[A-Z][a-z]+`)
}

func TestTeamCreature(t *testing.T) {
	testMatchRx(t, Team().Creature, `\w+`)
}

func TestTeamState(t *testing.T) {
	testMatchRx(t, Team().State, `[A-Z][a-z]+`)
}

func TestTeamStringer(t *testing.T) {
	rx := `[A-Z][a-z]+`
	res := fmt.Sprintf("%s", Team())
	if m, _ := regexp.MatchString(rx, res); !m {
		t.Errorf("expected %v to match %v", res, rx)
	}
}
