package faker

import (
	"fmt"
	"regexp"
	"testing"
)

func TestAvatarURL(t *testing.T) {
	rx := `http://robohash.org/\w+\.jpg\?size=100x200`
	res := Avatar().Url("jpg", 100, 200)
	if m, _ := regexp.MatchString(rx, res); !m {
		t.Errorf("expected %v to match %v", res, rx)
	}
}

func TestAvatarStringer(t *testing.T) {
	rx := `http://robohash.org/\w+\.png\?size=300x300`
	res := fmt.Sprintf("%s", Avatar())
	if m, _ := regexp.MatchString(rx, res); !m {
		t.Errorf("expected %v to match %v", res, rx)
	}
}
