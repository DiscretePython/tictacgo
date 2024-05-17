package tictacgo

import (
	"fmt"
)

const RESET = "\033[0m"
const RED = "\033[31m"
const BLUE = "\033[34m"

const GRID_STRING = "+---+---+---+\n| %s | %s | %s |\n+---+---+---+\n| %s | %s | %s |\n+---+---+---+\n| %s | %s | %s |\n+---+---+---+\n"
const GRID_HEIGHT_WIDTH = 3
const GRID_CELLS = GRID_HEIGHT_WIDTH * GRID_HEIGHT_WIDTH

var invalidPlayError = fmt.Errorf("Invalid move")

type Grid struct {
	plays [GRID_CELLS]int
}

func (g *Grid) MakePlay(position int, player int) error {
	if position < 0 || position > GRID_CELLS-1 {
		return invalidPlayError
	} else if player != 1 && player != 2 {
		return invalidPlayError
	} else if g.plays[position] != 0 {
		return invalidPlayError
	}

	g.plays[position] = player
	return nil
}

func (g Grid) Print() {
	vals := make([]string, GRID_CELLS)
	for idx, play := range g.plays {
		if play == 0 {
			vals[idx] = fmt.Sprintf("%d", idx+1)
		} else if play == 1 {
			vals[idx] = RED + "X" + RESET
		} else {
			vals[idx] = BLUE + "O" + RESET
		}
	}
	fmt.Printf(GRID_STRING, vals[0], vals[1], vals[2], vals[3],
		vals[4], vals[5], vals[6], vals[7], vals[8])
}

func (g Grid) HasWin() (bool, *int) {
	// Check rows/cols
	for i := 0; i < GRID_HEIGHT_WIDTH; i++ {
		if g.plays[i*GRID_HEIGHT_WIDTH] == g.plays[i*GRID_HEIGHT_WIDTH+1] &&
			g.plays[i*GRID_HEIGHT_WIDTH+1] == g.plays[i*GRID_HEIGHT_WIDTH+2] &&
			g.plays[i*GRID_HEIGHT_WIDTH+2] != 0 {
			return true, &g.plays[i*GRID_HEIGHT_WIDTH]
		}
		if g.plays[i] == g.plays[i+GRID_HEIGHT_WIDTH] &&
			g.plays[i+GRID_HEIGHT_WIDTH] == g.plays[i+2*GRID_HEIGHT_WIDTH] &&
			g.plays[i+2*GRID_HEIGHT_WIDTH] != 0 {
			return true, &g.plays[i]
		}
	}

	// Check diagonals
	// TODO: Make work with any size grid
	if g.plays[0] == g.plays[4] &&
		g.plays[4] == g.plays[8] &&
		g.plays[0] != 0 {
		return true, &g.plays[0]
	}
	if g.plays[2] == g.plays[4] &&
		g.plays[4] == g.plays[6] &&
		g.plays[6] != 0 {
		return true, &g.plays[2]
	}

	return false, nil
}

func (g *Grid) Reset() {
	g.plays = *new([9]int)
}
