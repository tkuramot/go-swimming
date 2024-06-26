package piscine

func strLen(str string) int {
	var l int
	for range str {
		l++
	}
	return l
}

func ConcatParams(args []string) string {
	var totalLen int
	for _, arg := range args {
		totalLen += strLen(arg) + 1
	}

	bytes := make([]byte, 0, totalLen)
	for _, arg := range args {
		bytes = append(bytes, arg...)
		bytes = append(bytes, '\n')
	}

	return string(bytes)
}
