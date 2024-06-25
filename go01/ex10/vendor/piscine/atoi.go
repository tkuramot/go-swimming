package piscine

func Atoi(s string) int {
	var i, sign int

	for _, c := range s {
		if sign == 0 && c == '-' {
			sign = -1
			continue
		}
		if sign == 0 && c == '+' {
			sign = 1
			continue
		}

		if c < '0' || c > '9' {
			return 0
		}
		i = i*10 + int(c-'0')
	}
	if sign == 0 {
		sign = 1
	}
	return i * sign
}
