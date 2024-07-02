package piscine

func Len[T any](a []T) int {
	var count int
	for range a {
		count++
	}
	return count
}
