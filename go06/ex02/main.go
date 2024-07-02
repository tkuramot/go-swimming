package main

import (
	"os"
	"piscine"
)

func main() {
	argc := piscine.Len(os.Args)
	if argc < 2 {
		piscine.PrintErrln("File name missing")
		os.Exit(1)
	}
	if argc > 2 {
		piscine.PrintErrln("Too many arguments")
		os.Exit(1)
	}

	piscine.DisplayFile(os.Args[1])
}
