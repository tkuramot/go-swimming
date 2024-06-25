package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := 0
	b := &a
	c := &b
	piscine.UltimatePointOne(&c)
	fmt.Println(a)
}
