package main

import (
	"os"
	"piscine"
)

func main() {
	if piscine.Len(os.Args) < 4 {
		piscine.Fprintln(os.Stderr, "Usage: "+os.Args[0]+" -c numBytes file1 [file2 ...]")
		os.Exit(1)
	}

	numBytes := piscine.Atoi(os.Args[2])
	if !piscine.Ztail(os.Args[3:], int64(numBytes)) {
		os.Exit(1)
	}
}
