package main

func canConstruct(ransomNote string, magazine string) bool {
	letters := make(map[byte]int, 26)
	for i := range 26 {
		letters[byte(i+97)] = 0
	}

	for i := range magazine {
		letters[magazine[i]]++
	}

	for i := range ransomNote {
		letters[ransomNote[i]]--
	}

	for i := range 26 {
		if letters[byte(i+97)] < 0 {
			return false
		}
	}
	return true
}
