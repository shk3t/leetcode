package problems

func isPalindrome(s string) bool {
	b, e := 0, len(s)-1
	vals := [2]byte{}

	for {
		if b >= e {
			return true
		}

		vals[0], vals[1] = s[b], s[e]
		skip := false

		for i := range vals {
			if vals[i] >= 65 && vals[i] <= 90 {
				vals[i] += 32
			} else if !(vals[i] >= 97 && vals[i] <= 122) && !(vals[i] >= 48 && vals[i] <= 57) {
				if i == 0 {
					b++
				} else {
					e--
				}
				skip = true
			}

		}

		if skip {
		} else if vals[0] == vals[1] {
			b++
			e--
		} else {
			return false
		}

	}
}
