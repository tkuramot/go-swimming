package piscine

func BasicAtoi(s string) int {
	var i int
	for _, v := range s {
		i = i*10 + int(v-'0')
	}
	return i
}
