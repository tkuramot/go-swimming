package piscine

func BasicJoin(strs []string) string {
	var result string
	for _, s := range strs {
		result += s
	}
	return result
}
