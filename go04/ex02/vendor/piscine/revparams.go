package piscine

import (
	"ft"
)

func sliceLen(slice []string) int {
	var l int
	for range slice {
		l++
	}
	return l
}

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func RevParams(args []string) {
	l := sliceLen(args)
	for i := l - 1; i >= 0; i-- {
		printStr(args[i])
		printStr("\n")
	}
}
