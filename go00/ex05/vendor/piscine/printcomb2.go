package piscine

import (
	"ft"
)

func printNumber(n int) {
	ft.PrintRune(rune(n/10) + '0')
	ft.PrintRune(rune(n%10) + '0')
}

func PrintComb2() {
	for i := 0; i <= 99; i++ {
		for j := i + 1; j <= 99; j++ {
			printNumber(i)
			ft.PrintRune(' ')
			printNumber(j)
			if i == 98 {
				ft.PrintRune('\n')
			} else {
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
		}
	}
}
