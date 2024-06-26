package piscine

import "ft"

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func PrintWordsTables(table []string) {
	for _, s := range table {
		printStr(s)
	}
}
