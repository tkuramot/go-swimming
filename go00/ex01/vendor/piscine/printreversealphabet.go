package piscine

import (
	"ft"
)

func PrintReverseAlphabet() {
	for i := 'z'; i >= 'a'; i-- {
		ft.PrintRune(i)
	}
	ft.PrintRune('\n')
}
