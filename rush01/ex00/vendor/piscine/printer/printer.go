package printer

import (
	"ft"
)

func PrintRune(r rune) {
	ft.PrintRune(r)
}

func Print(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func Println(s string) {
	Print(s)
	Print("\n")
}
