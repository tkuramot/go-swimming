package piscine

func isAlnum(c rune) bool {
	return (c >= '0' && c <= '9') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= 'a' && c <= 'z')
}

func Capitalize(s string) string {
	runes := []rune(s)
	wasAlnum := false
	for i, c := range runes {
		if !wasAlnum && c >= 'a' && c <= 'z' {
			runes[i] = c - 'a' + 'A'
		} else if wasAlnum && c >= 'A' && c <= 'Z' {
			runes[i] = c - 'A' + 'a'
		}
		wasAlnum = isAlnum(c)
	}
	return string(runes)
}
