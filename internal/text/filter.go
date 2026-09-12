// Package text contains functions performing operations on text.
package text

import (
	"strings"
)

// Filter defines with type we expect for our filters.
type Filter = map[string]interface{}

// ArrToFilter turns a simple array of strings into a filter we can work with.
func ArrToFilter(arr []string) Filter {
	filter := make(Filter)

	for _, item := range arr {
		filter[item] = struct{}{}
	}

	return filter
}

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
func ContainsBadWords(filter Filter, message string) map[string]uint64 {
	found := make(map[string]uint64)

	words := strings.Fields(strings.ToLower(message))

	for _, word := range words {
		word = normalize(strings.Trim(word, ".,!?;:\"'"))

		if _, exists := filter[word]; exists {
			found[word]++
		}
	}

	return found
}
