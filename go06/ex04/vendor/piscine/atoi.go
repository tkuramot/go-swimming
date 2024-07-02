package piscine

// returns -1 if the string is not a positive number
func Atoi(s string) int {
	var i int

	for _, c := range s {
		if c < '0' || c > '9' {
			return -1
		}
		i = i*10 + int(c-'0')
	}
	return i
}
