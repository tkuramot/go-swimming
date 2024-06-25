package piscine

import (
	"ft"
)

func sliceLen(s []string) int {
	var l int
	for range s {
		l++
	}
	return l
}

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func SortParams(args []string) {
	l := sliceLen(args)
	for i := 0; i < l; i++ {
		for j := 0; j < l-i-1; j++ {
			if args[j] > args[j+1] {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}

	for _, arg := range args {
		printStr(arg)
	}
}
