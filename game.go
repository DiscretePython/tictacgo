package tictacgo

import "fmt"

type Game struct {
	grid Grid
}

func New() Game {
	return Game{
		grid: Grid{},
	}
}

func (g *Game) Start() {
	g.grid.Reset()
	g.grid.Print()

	var position int
	player := 1

	fmt.Print("\n")

	for {
		ShowTurn(player)
		fmt.Print(ansiUpN(1) + ERASE_LINE + "Next Move: ")
		fmt.Scanf("%d%s", &position)

		if err := g.grid.MakePlay(position-1, player); err != nil {
			g.grid.SetError(&err)
			continue
		} else {
			g.grid.SetError(nil)
		}

		win, _ := g.grid.HasWin()
		if win {
			fmt.Println("Player", player, "wins!")
			break
		} else if g.grid.IsFull() {
			fmt.Println("Draw!")
			break
		}

		player = 3 - player
	}
}

func ShowTurn(player int) {
	mark := getPlayerMark(player)
	fmt.Print(
		SAVE_POSITION +
			ansiUpN(5) +
			ansiForwardN(16) +
			fmt.Sprintf("Turn: %s Player %d", mark, player) +
			RESTORE_POSITION,
	)
}
