package faker

import (
	"fmt"
	"regexp"
	"testing"
)

func TestLoremCharacters(t *testing.T) {
	nums := []int{3, 17, 293}
	for _, n := range nums {
		for i := 0; i < 3; i++ {
			res := Lorem().Characters(n)
			if len(res) != n {
				t.Errorf("expected %v to have len %d", res, n)
			}
		}
	}
}

func TestLoremWord(t *testing.T) {
	testMatchRx(t, Lorem().Word, `\w+`)
}

func TestLoremWords(t *testing.T) {
	nums := []int{3, 17, 293}
	for _, n := range nums {
		for i := 0; i < 3; i++ {
			res := Lorem().Words(n)
			if len(res) != n {
				t.Errorf("expected %v to be %d words", res, n)
			}
		}
	}
}

func TestLoremSentence(t *testing.T) {
	rx := `[A-Z]\w*\s\w+\s\w+\.`
	for i := 0; i < 10; i++ {
		res := Lorem().Sentence(3)
		if m, _ := regexp.MatchString(rx, res); !m {
			t.Errorf("expected %v to match %v", res, rx)
		}
	}

}

func TestLoremSentences(t *testing.T) {
	nums := []int{3, 17}
	for _, n := range nums {
		for i := 0; i < 3; i++ {
			res := Lorem().Sentences(n)
			if len(res) != n {
				t.Errorf("expected %v to have len %d", res, n)
			}
		}
	}
}

func TestLoremParagraph(t *testing.T) {
	rx := `[^\.]+\.[^\.]+\.[^\.]+\.`
	for i := 0; i < 10; i++ {
		res := Lorem().Paragraph(3)
		if m, _ := regexp.MatchString(rx, res); !m {
			t.Errorf("expected %v to match %v", res, rx)
		}
	}
}

func TestLoremParagraphs(t *testing.T) {
	nums := []int{3, 17}
	for _, n := range nums {
		for i := 0; i < 3; i++ {
			res := Lorem().Paragraphs(n)
			if len(res) != n {
				t.Errorf("expected %v to have len %d", res, n)
			}
		}
	}
}

func TestLoremStringer(t *testing.T) {
	rx := `^[A-Z]\w*\s[^\.]+\.$`
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("%s", Lorem())
		if m, _ := regexp.MatchString(rx, res); !m {
			t.Errorf("expected %v to match %v", res, rx)
		}
	}
}
