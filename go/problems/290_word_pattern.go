package problems

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Fields(s)

	if len(pattern) != len(words) {
		return false
	}

	pwMapping := [26]string{}
	mappedWords := make(map[string]struct{}, 26)

	for i := range pattern {
		l, word := pattern[i], words[i]
		pMapping := pwMapping[l-97]
		_, wordMapped := mappedWords[word]
		if (pMapping == "" && !wordMapped) || pMapping == word {
			pwMapping[l-97] = word
			mappedWords[word] = struct{}{}
		} else {
			return false
		}
	}

	return true
}
