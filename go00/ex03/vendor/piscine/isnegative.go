package piscine

import (
	"ft"
)

func IsNegative(n int) {
	if n < 0 {
		ft.PrintRune('T')
	} else {
		ft.PrintRune('F')
	}
	ft.PrintRune('\n')
}
