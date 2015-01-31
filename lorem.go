package faker

import (
	"math/rand"
	"strings"
)

type Lorem struct{}

// Example:
//	Lorem{}.Character() // c
func (l Lorem) Character() string {
	return l.Characters(1)
}

// Example:
//	Lorem{}.Characters(17) // wqFyJIrXYfVP7cL9M
func (l Lorem) Characters(num int) string {
	perm := rand.Perm(len(DigitsAndLetters) * (num/len(DigitsAndLetters) + 1))
	res := make([]byte, num)
	for i := range res {
		res[i] = DigitsAndLetters[perm[i]%len(DigitsAndLetters)]
	}
	return string(res)
}

// Example:
//	Lorem{}.Word() // veritatis
func (l Lorem) Word() string {
	return Fetch("lorem.words")
}

// Example:
//	Lorem{}.Words(3) // [omnis libero neque]
func (l Lorem) Words(num int) []string {
	words := valueAt("lorem.words").([]string)
	perm := rand.Perm(len(words) * (num/len(words) + 1))
	res := make([]string, num)
	for i := range res {
		res[i] = words[perm[i]%len(words)]
	}
	return res
}

// Sentence returns random capitalized sentence of "words" words length.
// Example:
//	Lorem{}.Sentence(3) // Necessitatibus cum autem.
func (l Lorem) Sentence(words int) string {
	s := strings.Join(l.Words(words), " ")
	return strings.ToTitle(s[:1]) + s[1:] + "."
}

// Sentences returns a slice of "num" sentences, 3 to 11 words each.
func (l Lorem) Sentences(num int) []string {
	res := make([]string, num)
	for i := range res {
		res[i] = l.Sentence(rand.Intn(9) + 3) // 3 to 11 words
	}
	return res
}

// Paragraph returns a random text of "sentences" sentences length.
func (l Lorem) Paragraph(sentences int) string {
	return strings.Join(l.Sentences(sentences), " ")
}

// Paragraphs returns a slice of "num" paragraphs, 3 to 11 sentences each.
func (l Lorem) Paragraphs(num int) []string {
	res := make([]string, num)
	for i := range res {
		res[i] = l.Paragraph(rand.Intn(9) + 3) // 3 to 11 sentences
	}
	return res
}

// String returns a random sentence 3 to 11 words in length.
func (l Lorem) String() string {
	return l.Sentence(rand.Intn(9) + 3) // 3 to 11 words
}
