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

	for {
		fmt.Print("Next Move: ")
		fmt.Scanf("%d%s", &position)
		if err := g.grid.MakePlay(position-1, player); err != nil {
			fmt.Println(err)
			continue
		}

		g.grid.Print()
		win, _ := g.grid.HasWin()
		if win {
			fmt.Println("Player", player, "wins!")
			break
		}

		player = 3 - player
	}
}
