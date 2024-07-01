package grids

import (
	"piscine/printer"
	"piscine/unicode"
)

const (
	size    = 5
	maxSize = 10
	Black   = 'B'
	White   = '.'
)

var directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

type Grid struct {
	Cells [maxSize][maxSize]rune
	Size  int
}

func NewGrid(args []string) *Grid {
	grid := &Grid{}
	for i, row := range args {
		for j, cell := range row {
			grid.Cells[i][j] = cell
		}
	}
	grid.Size = len(args) // args must be square
	return grid
}

func (g *Grid) Print() {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			printer.PrintRune(g.Cells[i][j])
		}
		printer.Println("")
	}
}

func (g *Grid) PlaceBlack(x, y int) {
	g.Cells[x][y] = Black
}

func (g *Grid) RemoveBlack(x, y int) {
	g.Cells[x][y] = White
}

func (g *Grid) IsWhite(x, y int) bool {
	return g.Cells[x][y] == White
}

func (g *Grid) IsBlack(x, y int) bool {
	return g.Cells[x][y] == Black
}

func (g *Grid) CanPlaceBlack(x, y int) bool {
	for _, d := range directions {
		nx, ny := x+d[0], y+d[1]
		if nx >= 0 && nx < size && ny >= 0 && ny < size && g.IsBlack(nx, ny) {
			return false
		}
	}
	return true
}

// white space in this context is a literal white space and numbers
func (g *Grid) IsWhiteSpaceVisible(x, y int) bool {
	count := 0
	for _, d := range directions {
		for i := 1; i < size; i++ {
			nx, ny := x+i*d[0], y+i*d[1]
			if nx < 0 || nx >= size || ny < 0 || ny >= size || g.IsBlack(nx, ny) {
				break
			}
			count++
		}
	}
	return count == unicode.RuneToInt(g.Cells[x][y])-1
}

func (g *Grid) floodFill(x, y int, visited *[maxSize][maxSize]bool) {
	if x < 0 || x >= g.Size || y < 0 || y >= g.Size || g.IsBlack(x, y) || visited[x][y] {
		return
	}
	visited[x][y] = true
	for _, d := range directions {
		g.floodFill(x+d[0], y+d[1], visited)
	}
}

func (g *Grid) isAllWhiteConnected() bool {
	var visited [maxSize][maxSize]bool
	started := false

	for i := 0; i < g.Size; i++ {
		for j := 0; j < g.Size; j++ {
			if g.IsWhite(i, j) {
				if !started {
					g.floodFill(i, j, &visited)
					started = true
				} else if !visited[i][j] {
					return false
				}
			}
		}
	}
	return true
}

func (g *Grid) IsSolved() bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if unicode.IsDigit(g.Cells[i][j]) && !g.IsWhiteSpaceVisible(i, j) {
				return false
			}
		}
	}
	return g.isAllWhiteConnected()
}
