package piscine

func IsPrintable(str string) bool {
	for _, r := range str {
		if r < 32 || r > 126 {
			return false
		}
	}
	return true
}
