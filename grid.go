package tictacgo

import (
	"fmt"
)

const GRID_STRING = "+---+---+---+\n| %s | %s | %s |\n+---+---+---+\n| %s | %s | %s |\n+---+---+---+\n| %s | %s | %s |\n+---+---+---+\n"
const GRID_HEIGHT_WIDTH = 3
const GRID_CELLS = GRID_HEIGHT_WIDTH * GRID_HEIGHT_WIDTH

var invalidPlayError = fmt.Errorf("Invalid move")

type Grid struct {
	plays [GRID_CELLS]int
	err   *error
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

	col := (position) % GRID_HEIGHT_WIDTH
	row := (position) / GRID_HEIGHT_WIDTH
	fmt.Print(
		SAVE_POSITION +
			ansiUpN(2*(GRID_HEIGHT_WIDTH-row)+1) +
			ansiForwardN(4*(col)+2) +
			getPlayerMark(player) +
			RESTORE_POSITION,
	)
	return nil
}

func getPlayerMark(player int) string {
	if player == 1 {
		return RED + "X" + RESET
	} else {
		return BLUE + "O" + RESET
	}
}

func (g Grid) Print() {
	vals := make([]string, GRID_CELLS)
	for idx, play := range g.plays {
		if play == 0 {
			vals[idx] = fmt.Sprintf("%d", idx+1)
		} else {
			vals[idx] = getPlayerMark(play)
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

func (g Grid) IsFull() bool {
	for _, play := range g.plays {
		if play == 0 {
			return false
		}
	}
	return true
}

func (g *Grid) Reset() {
	g.plays = *new([9]int)
}

func (g *Grid) SetError(err *error) {
	g.err = err

	if err != nil {
		fmt.Print(
			SAVE_POSITION+ansiUpN(2)+ansiForwardN(16),
			*err,
			RESTORE_POSITION,
		)
	} else {
		fmt.Print(
			SAVE_POSITION +
				ansiUpN(2) + ansiForwardN(16) +
				ERASE_LINE_TO_END + RESTORE_POSITION,
		)
	}
}
