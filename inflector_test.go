package inflect

import (
	"testing"
)

func TestInflect(t *testing.T) {
	tests := []struct {
		Kebab   string
		Public  string
		Private string
	}{
		{
			Kebab:   "word",
			Public:  "Word",
			Private: "word",
		},
		{
			Kebab:   "word-pair",
			Public:  "WordPair",
			Private: "wordPair",
		},
		// Known acronym.
		{
			Kebab:   "cpu",
			Public:  "CPU",
			Private: "cpu",
		},
	}
	check := func(f func(string) string, name, input, expected string) {
		actual := f(input)
		if expected != actual {
			t.Errorf("expected %s(%q) to equal %q, but got %q", name, input, expected, actual)
		}
	}
	for _, test := range tests {
		check(KebabToPublic, "KebabToPublic", test.Kebab, test.Public)
		check(KebabToPrivate, "KebabToPrivate", test.Kebab, test.Private)
	}
}
