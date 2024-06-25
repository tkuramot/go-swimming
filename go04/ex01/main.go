package main

import (
	"os"
	"piscine"
)

func main() {
	piscine.PrintParams(os.Args[1:])
}
