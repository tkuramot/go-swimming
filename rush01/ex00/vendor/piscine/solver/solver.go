package solver

import (
	"piscine/grids"
)

func Solve(g *grids.Grid, x, y int) bool {
	if x == g.Size {
		x = 0
		y++
		if y == g.Size {
			return g.IsSolved()
		}
	}

	if g.IsWhite(x, y) {
		if g.CanPlaceBlack(x, y) {
			g.PlaceBlack(x, y)
			if Solve(g, x+1, y) {
				return true
			}
			g.RemoveBlack(x, y)
		}
		if Solve(g, x+1, y) {
			return true
		}
	} else {
		if Solve(g, x+1, y) {
			return true
		}
	}

	return false
}
