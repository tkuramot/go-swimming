package main

import (
	"os"
	"piscine/grids"
	"piscine/printer"
	"piscine/solver"
)

func main() {
	// if len(os.Args) != 6 {
	// 	printer.Println("Invalid number of arguments")
	// 	os.Exit(1)
	// }
	args := os.Args[1:]
	grid := grids.NewGrid(args)
	if grid == nil {
		printer.Println("Invalid grid") // TODO more specific error
		os.Exit(1)
	}
	if solver.Solve(grid, 0, 0) {
		grid.Print()
	} else {
		printer.Println("No solution")
	}
}
