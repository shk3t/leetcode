package main

func isValid(s string) bool {
	st := []rune{}

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			st = append(st, v)
		} else if len(st) > 0 && (
			(v == ')' && st[len(st)-1] == '(') ||
			(v == ']' && st[len(st)-1] == '[') ||
			(v == '}' && st[len(st)-1] == '{')) {
			st = st[:len(st)-1]
		} else {
			return false
		}
	}

	return len(st) == 0
}
