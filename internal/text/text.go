// Package text contains functions performing operations on text.
package text

import (
	"strings"
)

func normalize(s string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case '!', '1', '|':
			return 'i'
		case '$', '5':
			return 's'
		case '@':
			return 'a'
		case '0':
			return 'o'
		case '3':
			return 'e'
		default:
			return r
		}
	}, s)
}

// ContainsBadWords checks if a string contains any bad words from a filter array and returns the found bad words with count.
func ContainsBadWords(filter map[string]bool, message string) map[string]int {
	found := make(map[string]int)

	words := strings.Fields(strings.ToLower(message))

	for _, word := range words {
		word = normalize(strings.Trim(word, ".,!?;:\"'"))

		if filter[word] {
			found[word]++
		}
	}

	return found
}
