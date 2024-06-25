package piscine

func strLen(s string) int {
	var l int
	for range s {
		l++
	}
	return l
}

func index(s string, toFind rune) int {
	for i, c := range s {
		if c == rune(toFind) {
			return i
		}
	}
	return -1
}

func isBaseValid(base string) bool {
	l := strLen(base)
	if l < 2 {
		return false
	}
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if base[i] == '+' || base[i] == '-' {
				return false
			}
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}

func AtoiBase(s string, base string) int {
	if !isBaseValid(base) {
		return 0
	}
	var n int
	bl := strLen(base)
	for _, r := range s {
		n = n*bl + index(base, r)
	}
	return n
}
