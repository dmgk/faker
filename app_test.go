package faker

import (
	"fmt"
	"regexp"
	"testing"
)

func TestAppName(t *testing.T) {
	testMatchRx(t, App().Name, `[A-Z][-\w]+`)
}

func TestAppVersion(t *testing.T) {
	testMatchRx(t, App().Version, `\d+\.\d+(\.\d+)?`)
}

func TestAppAuthor(t *testing.T) {
	testMatchRx(t, App().Author, `[A-Z]\w+`)
}

func TestAppStringer(t *testing.T) {
	rx := `[A-Z][-\w]+`
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("%v", App())
		if m, _ := regexp.MatchString(rx, res); !m {
			t.Errorf("expected %v to match %v", res, rx)
		}
	}
}
