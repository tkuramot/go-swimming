package piscine

import (
	"ft"
	"os"
)

func printStr(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
}

func strLen(s string) int {
	var l int
	for range s {
		l++
	}
	return l
}

func extractProgramName(s string) string {
  l := strLen(s)
	for i := l - 1; i >= 0; i-- {
		if s[i] == '/' {
			return s[i+1:]
		}
	}
	return s
}

func PrintProgramName() {
	name := extractProgramName(os.Args[0])
	printStr(name)
}
