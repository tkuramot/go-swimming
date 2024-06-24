package piscine

import (
	"ft"
)

func PrintComb() {
	for i := '0'; i <= '9'; i++ {
		for j := i + 1; j <= '9'; j++ {
			for k := j + 1; k <= '9'; k++ {
				ft.PrintRune(i)
				ft.PrintRune(j)
				ft.PrintRune(k)
				if i == '7' {
					ft.PrintRune('\n')
				} else {
					ft.PrintRune(',')
					ft.PrintRune(' ')
				}
			}
		}
	}
}
