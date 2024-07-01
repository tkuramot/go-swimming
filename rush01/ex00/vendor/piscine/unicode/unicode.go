package unicode

func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func RuneToInt(r rune) int {
	return int(r - '0')
}
