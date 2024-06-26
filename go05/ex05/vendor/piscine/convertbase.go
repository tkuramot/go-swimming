package piscine

import (
	"ft"
)

func strLen(s string) int {
	var l int
	for range s {
		l++
	}
	return l
}

func printStr(str string) {
	for _, r := range str {
		ft.PrintRune(r)
	}
}

func nRune(s string, n int) rune {
	for i, r := range s {
		if i == n-1 {
			return r
		}
	}
	return 0
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

// string s is guaranteed to be valid according to the subject
func atoiBase(s string, base string) int {
	if !isBaseValid(base) {
		return 0
	}
	var n int
	bl := strLen(base)
	for _, r := range s {
		n = n*bl + index(base, string(r))
	}
	return n
}

func itoaBase(nbr int, base string) string {
	if !isBaseValid(base) {
		return "NV"
	}

	if nbr == 0 {
		return ""
	}

	l := strLen(base)
	if nbr < 0 {
		return "-" + itoaBase(-nbr, base)
	}
	return itoaBase(nbr/l, base) + string(nRune(base, nbr%l+1))
}

// bases are guaranteed to be valid according to the subject
// does not handle negative numbers
func ConvertBase(nbr, baseFrom, baseTo string) string {
	i := atoiBase(nbr, baseFrom)
	return itoaBase(i, baseTo)
}
