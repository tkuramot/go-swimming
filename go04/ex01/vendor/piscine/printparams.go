package piscine

import (
	"ft"
)

func printStr(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
}

func PrintParams(args []string) {
	for _, arg := range args {
		printStr(arg)
		printStr("\n")
	}
}
