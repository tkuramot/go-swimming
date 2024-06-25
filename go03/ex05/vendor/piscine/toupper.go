package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	for i := range runes {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			runes[i] = runes[i] - 'a' + 'A'
		}
	}
	return string(runes)
}
