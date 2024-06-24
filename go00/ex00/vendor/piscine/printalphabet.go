package piscine

import (
	"ft"
)

func PrintAlphabet() {
	for i := 'a'; i <= 'z'; i++ {
		ft.PrintRune(i)
	}
	ft.PrintRune('\n')
}
