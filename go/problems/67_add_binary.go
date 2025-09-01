package problems

// Solution 1

func addBinary(a string, b string) string {
	sum := []rune{}
	i, j := len(a)-1, len(b)-1

	nextBit := '0'
	for i >= 0 || j >= 0 {
		if i < 0 {
			if b[j] == '0' {
				sum, nextBit = append(sum, nextBit), '0'
			} else {
				sum = set1Bit(sum, nextBit)
			}
		} else if j < 0 {
			if a[i] == '0' {
				sum, nextBit = append(sum, nextBit), '0'
			} else {
				sum = set1Bit(sum, nextBit)
			}
		} else if a[i] == '1' && b[j] == '1' {
			sum, nextBit = append(sum, nextBit), '1'
		} else if a[i] == '1' || b[j] == '1' {
			sum = set1Bit(sum, nextBit)
		} else {
			sum, nextBit = append(sum, nextBit), '0'
		}
		i--
		j--
	}

	if nextBit == '1' {
		sum = append(sum, nextBit)
	}

	reverse(sum)
	return string(sum)
}

func set1Bit(sum []rune, nextBit rune) []rune {
	if nextBit == '1' {
		return append(sum, '0')
	} else {
		return append(sum, '1')
	}
}

func reverse[T comparable](arr []T) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

// Solution 2

func addBinary2(a string, b string) string {
	sumRep := make([]rune, max(len(a), len(b))+1)

	sum := 0
	i, j := len(a)-1, len(b)-1

	for i >= 0 || j >= 0 {
		if i >= 0 && a[i] == '1' {
			sum++
		}
		if j >= 0 && b[j] == '1' {
			sum++
		}
		i--
		j--

		sumRep[max(i, j)+2] = rune('0' + sum%2)
		sum /= 2
	}

	if sum != 0 {
		sumRep[0] = rune('0' + sum%2)
		return string(sumRep)
	} else {
		return string(sumRep[1:])
	}
}
