package piscine

import (
	"ft"
)

func printSlice(arr []int) {
	for _, v := range arr {
		ft.PrintRune(rune(v) + '0')
	}
}

func isConsecutiveSlice(slice []int) bool {
	if len(slice) < 2 {
		return true
	}

	sz := len(slice)
	for i := 1; i < sz; i++ {
		if slice[i]-slice[i-1] != 1 {
			return false
		}
	}

	return true
}

func printCombNHelper(comb []int, idx int, start int, comb_n int) {
	if idx == comb_n {
		if comb[0] != 0 || !isConsecutiveSlice(comb) {
			ft.PrintRune(',')
			ft.PrintRune(' ')
		}
		printSlice(comb)
		return
	}

	for i := start; i <= 9; i++ {
		comb[idx] = i
		printCombNHelper(comb, idx+1, i+1, comb_n)
	}
}

func PrintCombN(n int) {
	if n < 1 || n > 9 {
		return
	}
	comb := make([]int, n)
	printCombNHelper(comb, 0, 0, n)
	ft.PrintRune('\n')
}
