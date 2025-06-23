package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letterCount := [26]int{}

	for i := range s {
		letterCount[s[i]-97]++
		letterCount[t[i]-97]--
	}

	for _, count := range letterCount {
		if count != 0 {
			return false
		}
	}

	return true
}
