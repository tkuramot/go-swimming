package piscine

import (
	"os"
)

func Fprint(f *os.File, s string) {
	f.WriteString(s)
}

func Fprintln(f *os.File, s string) {
	Fprint(f, s)
	Fprint(f, "\n")
}

func Print(s string) {
	Fprint(os.Stdout, s)
}

func Println(s string) {
	Fprintln(os.Stdout, s)
}

func PrintErr(s string) {
	Fprint(os.Stderr, "ERROR: ")
	Fprint(os.Stderr, s)
}

func PrintErrln(s string) {
	Fprint(os.Stderr, "ERROR: ")
	Fprintln(os.Stderr, s)
}
