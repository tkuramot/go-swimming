package piscine

import (
	"os"
)

func Cat(args []string) {
	if Len(args) == 0 {
		ReadAndPrintFile(os.Stdin)
		return
	}

	for _, arg := range args {
		DisplayFile(arg)
	}
}

