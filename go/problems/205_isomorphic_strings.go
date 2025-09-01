package problems

func isIsomorphic(s string, t string) bool {
	mapping := [128]byte{}
	mapped := [128]bool{}

	for i := 0; i < len(s); i++ {
		sMapping := mapping[s[i]]
		tMapped := mapped[t[i]]
		if (sMapping == 0 && !tMapped) || sMapping == t[i] {
			mapping[s[i]] = t[i]
			mapped[t[i]] = true
		} else {
			return false
		}
	}

	return true
}
