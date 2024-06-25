package piscine

import (
	"ft"
)

func PrintStr(s string) {
	for _, v := range s {
		ft.PrintRune(v)
	}
}
