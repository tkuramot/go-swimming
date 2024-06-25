package piscine

func sliceLen(s []int) int {
	var l int
	for range s {
		l++
	}
	return l
}

func SortIntegerTable(table []int) {
	l := sliceLen(table)
	for i := 0; i < l; i++ {
		for j := 0; j < l-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}
