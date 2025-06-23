package main

func lengthOfLastWord(s string) int {
	l := 0
	prev_space := false

	for i := 0; i < len(s); i++ {
		if s[i] == 32 {
			prev_space = true
		} else if prev_space {
			l = 1
			prev_space = false
		} else {
			l++
		}
	}

	return l
}
