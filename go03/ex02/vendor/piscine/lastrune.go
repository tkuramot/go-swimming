package piscine

func LastRune(s string) rune {
	var l int
	for range s {
		l++
	}
	runes := []rune(s)
	return runes[l-1]
}
