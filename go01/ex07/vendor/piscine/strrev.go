package piscine

func strLen(s string) int {
	var l int
	for range s {
		l++
	}
	return l
}

func StrRev(s string) string {
	r := []rune(s)
	l := strLen(s)
	for i := 0; i < l/2; i++ {
		r[i], r[l-i-1] = r[l-i-1], r[i]
	}
	return string(r)
}
