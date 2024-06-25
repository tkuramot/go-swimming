package piscine

import (
	"ft"
)

func printStr(str string) {
	for _, r := range str {
		ft.PrintRune(r)
	}
}

func strLen(str string) int {
	var l int
	for range str {
		l++
	}
	return l
}

func nRune(s string, n int) rune {
	for i, r := range s {
		if i == n-1 {
			return r
		}
	}
	return 0
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

func PrintNbrBase(nbr int, base string) {
	if !isBaseValid(base) {
		printStr("NV")
		return
	}

	if nbr == 0 {
		return
	}

	l := strLen(base)
	if nbr < 0 {
		ft.PrintRune('-')
		nbr = -nbr
	}
	PrintNbrBase(nbr/l, base)
	ft.PrintRune(nRune(base, nbr%l+1))
}
