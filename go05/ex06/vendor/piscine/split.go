package piscine

func strLen(s string) int {
	var l int
	for range s {
		l++
	}
	return l
}

func count(s, substr string) int {
	if strLen(substr) == 0 {
		return strLen(s)
	}

	var n int
	for {
		i := index(s, substr)
		if i == -1 {
			return n
		}
		n++
		s = s[i+strLen(substr):]
	}
}

func index(s string, toFind string) int {
	for i, c := range s {
		if c == rune(toFind[0]) {
			if s[i:i+strLen(toFind)] == toFind {
				return i
			}
		}
	}
	return -1
}

func explode(s string, n int) []string {
	a := make([]string, n)
	for i, c := range s {
		a[i] = string(c)
	}
	return a
}

func Split(str, charset string) []string {
	n := count(str, charset) + 1
	if charset == "" {
		return explode(str, n - 1)
	}

	a := make([]string, n)
	n--
	var i int
	for i < n {
		m := index(str, charset)
		if m < 0 {
			break
		}
		a[i] = str[:m]
		str = str[m+strLen(charset):]
		i++
	}
	a[i] = str
	return a
}
