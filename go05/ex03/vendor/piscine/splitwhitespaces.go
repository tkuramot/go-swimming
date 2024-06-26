package piscine

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\r': 1, ' ': 1}

func strLen(s string) int {
	var l int
	for range s {
		l++
	}
	return l
}

func SplitWhiteSpaces(str string) []string {
	n := 0
	wasSpace := 1

	l := strLen(str)
	for i := 0; i < l; i++ {
		r := str[i]
		isSpace := int(asciiSpace[r])
		n += wasSpace & ^isSpace
		wasSpace = isSpace
	}

	a := make([]string, n)
	na := 0
	fieldStart := 0
	var i int
	for i < l && asciiSpace[str[i]] != 0 {
		i++
	}
	fieldStart = i
	for i < l {
		if asciiSpace[str[i]] == 0 {
			i++
			continue
		}
		a[na] = str[fieldStart:i]
		na++
		i++
		for i < l && asciiSpace[str[i]] != 0 {
			i++
		}
		fieldStart = i
	}
	if fieldStart < l {
		a[na] = str[fieldStart:]
	}
	return a
}
