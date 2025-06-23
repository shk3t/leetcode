package main

func romanToInt(s string) int {
	prev := ' '
	sum := 0

	for _, v := range s {
		switch v {
		case 'I':
			sum += 1
		case 'V':
			sum += 5
			if prev == 'I' {
				sum -= 2
			}
		case 'X':
			sum += 10
			if prev == 'I' {
				sum -= 2
			}
		case 'L':
			sum += 50
			if prev == 'X' {
				sum -= 20
			}
		case 'C':
			sum += 100
			if prev == 'X' {
				sum -= 20
			}
		case 'D':
			sum += 500
			if prev == 'C' {
				sum -= 200
			}
		case 'M':
			sum += 1000
			if prev == 'C' {
				sum -= 200
			}
		}
		prev = v
	}

	return sum
}
