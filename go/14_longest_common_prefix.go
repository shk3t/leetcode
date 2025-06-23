package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := []byte{}

	for j := 0; ; j++ {
		if j >= len(strs[0]) {
			return string(prefix)
		}

		b := strs[0][j]
		for i := 1; i < len(strs); i++ {
			if j >= len(strs[i]) || strs[i][j] != b {
				return string(prefix)
			}
		}

		prefix = append(prefix, b)
	}
}
