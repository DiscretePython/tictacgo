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

	fmt.Print("Next Move: ")
	for {
		fmt.Scanf("%d%s", &position)
		if err := g.grid.MakePlay(position-1, player); err != nil {
			g.grid.SetError(&err)
			fmt.Print(ansiUpN(1) + ERASE_LINE + "Next Move: ")
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
		fmt.Print(ansiUpN(1) + ERASE_LINE + "Next Move: ")
	}
}
