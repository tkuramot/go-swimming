package piscine

import (
	"fmt"
)

const N = 8

var directions = [8][2]int{
	{1, 0}, {0, 1}, {-1, 0}, {0, -1}, // Up, Right, Down, Left
	{1, 1}, {1, -1}, {-1, 1}, {-1, -1}, // Diagonals
}

func simulate(board *[N][N]int, row int, column int, value int) {
	board[row][column] += value
	for _, dir := range directions {
		for d := 1; d < N; d++ {
			newRow, newCol := row+d*dir[0], column+d*dir[1]
			if newRow >= 0 && newRow < N && newCol >= 0 && newCol < N {
				board[newRow][newCol] += value
			} else {
				break
			}
		}
	}
}

func eightQueensHelper(queens *[N]int, board [N][N]int, column int) {
	if column == N {
		fmt.Println(queens)
		return
	}

	for row := 0; row < N; row++ {
		if board[row][column] == 0 {
			queens[column] = row
			simulate(&board, row, column, 1)
			eightQueensHelper(queens, board, column+1)
			simulate(&board, row, column, -1)
		}
	}
}

func EightQueens() {
	queens := [N]int{}
	board := [N][N]int{}

	eightQueensHelper(&queens, board, 0)
}
