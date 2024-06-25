package piscine

func strLen(s string) int {
	var l int
	for range s {
		l++
	}
	return l
}

func Index(s string, toFind string) int {
	for i, c := range s {
		if c == rune(toFind[0]) {
			if s[i:i+strLen(toFind)] == toFind {
				return i
			}
		}
	}
	return -1
}
