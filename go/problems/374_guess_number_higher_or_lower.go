package problems

func guess(num int) int {
	return 0
}

func guessNumber(n int) int {
	l, r := 1, n

	for {
		pick := (l + r) / 2
		switch guess(pick) {
		case -1:
			r = pick
		case 1:
			l = pick + 1
		default:
			return pick
		}
	}
}
