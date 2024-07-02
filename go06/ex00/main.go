package main

import (
	"ft"
	"os"
)

type boolean int

const (
	yes boolean = 1
	no  boolean = 0
)

const (
	EvenMsg = "I have an even number of arguments"
	OddMsg  = "I have an odd number of arguments"
)

var lengthOfArg int = sliceLen(os.Args)

func sliceLen(s []string) int {
	var i int
	for range s {
		i++
	}
	return i
}

func even(nbr int) boolean {
	if nbr%2 == 0 {
		return yes
	}
	return no
}

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func isEven(nbr int) boolean {
	if even(nbr) == 1 {
		return yes
	} else {
		return no
	}
}

func main() {
	if isEven(lengthOfArg) == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
