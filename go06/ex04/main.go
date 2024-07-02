package main

import (
	"os"
	"piscine"
)

func main() {
	if piscine.Len(os.Args) < 4 || os.Args[1] != "-c" {
		piscine.Fprintln(os.Stderr, "Usage: "+os.Args[0]+" -c numBytes file1 [file2 ...]")
		os.Exit(1)
	}

	numBytes := piscine.Atoi(os.Args[2])
	if numBytes < 0 {
		piscine.Fprintln(os.Stderr, os.Args[0]+": invalid number of bytes: '"+os.Args[2]+"'")
		os.Exit(1)
	}
	if !piscine.Ztail(os.Args[3:], int64(numBytes)) {
		os.Exit(1)
	}
}
