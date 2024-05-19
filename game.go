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
	g.grid.PrintNewGrid()

	var position int
	player := 1

	fmt.Print("\n")

	for {
		ShowTurn(player)
		fmt.Print(ansiUpN(1) + ERASE_LINE + "Next Move: ")

		_, err := fmt.Scanf("%d\n", &position)
		if err != nil && err.Error() != "unexpected newline" {
			for {
				_, err := fmt.Scanln()
				if err == nil {
					break
				}
			}
		} else if err != nil {
			continue
		}

		if err := g.grid.MakePlay(position-1, player); err != nil {
			g.grid.ShowError()
			continue
		} else if err := g.grid.PrintPlay(position-1, player); err != nil {
			g.grid.ShowError()
			continue
		} else if g.grid.Err != nil {
			g.grid.ClearError()
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
