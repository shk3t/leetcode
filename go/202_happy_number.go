package main

func isHappy(n int) bool {
	history := map[int]struct{}{n: struct{}{}}
	digits := [10]int{}

	for {
		temp := n
		i := 0
		for temp > 0 {
			digits[i] = temp % 10
			temp /= 10
			i++
		}

		n = 0
		for j := 0; j < i; j++ {
			n += digits[j] * digits[j]
		}

		if _, ok := history[n]; ok {
			return n == 1
		}
		history[n] = struct{}{}
	}
}
