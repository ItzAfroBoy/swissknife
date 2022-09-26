package swissknife

import "strings"

func Lines(s string) int {
	return strings.Count(s, "\n")
}

func Longest(s string, delim string) (int, string) {
	words := strings.Split(s, delim)
	length := 0
	var wd string

	for _, word := range words {
		if len(word) > length {
			length = len(word)
			wd = word
		}
	}

	return length, wd
}
