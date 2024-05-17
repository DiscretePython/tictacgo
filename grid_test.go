package tictacgo_test

import (
	"testing"
	"tictacgo"
)

func TestGrid(t *testing.T) {
	t.Run("Win condition works row", func(t *testing.T) {
		grid := tictacgo.Grid{}

		grid.MakePlay(0, 1)
		grid.MakePlay(1, 1)
		grid.MakePlay(2, 1)

		win, _ := grid.HasWin()
		if !win {
			t.Errorf("Expected win to be true")
		}
	})

	t.Run("Win condition works col", func(t *testing.T) {
		grid := tictacgo.Grid{}

		grid.MakePlay(0, 1)
		grid.MakePlay(3, 1)
		grid.MakePlay(6, 1)

		win, _ := grid.HasWin()
		if !win {
			t.Errorf("Expected win to be true")
		}
	})

	t.Run("Win condition works diag", func(t *testing.T) {
		grid := tictacgo.Grid{}

		grid.MakePlay(0, 1)
		grid.MakePlay(4, 1)
		grid.MakePlay(8, 1)

		win, _ := grid.HasWin()
		if !win {
			t.Errorf("Expected win to be true")
		}
	})

	t.Run("Win condition correct player 1 won", func(t *testing.T) {
		grid := tictacgo.Grid{}

		grid.MakePlay(0, 1)
		grid.MakePlay(1, 1)
		grid.MakePlay(2, 1)

		_, player := grid.HasWin()
		if *player != 1 {
			t.Errorf("Expected player to be 1")
		}
	})

	t.Run("Win condition correct player 2 won", func(t *testing.T) {
		grid := tictacgo.Grid{}

		grid.MakePlay(0, 2)
		grid.MakePlay(1, 2)
		grid.MakePlay(2, 2)

		_, player := grid.HasWin()
		if *player != 2 {
			t.Errorf("Expected player to be 2")
		}
	})

	t.Run("Win condition no win", func(t *testing.T) {
		grid := tictacgo.Grid{}

		grid.MakePlay(0, 1)
		grid.MakePlay(1, 2)
		grid.MakePlay(2, 1)

		win, player := grid.HasWin()
		if win {
			t.Errorf("Expected win to be false")
		} else if player != nil {
			t.Errorf("Expected player to be nil")
		}
	})

	t.Run("Reset clears win", func(t *testing.T) {
		grid := tictacgo.Grid{}

		grid.MakePlay(0, 2)
		grid.MakePlay(1, 2)
		grid.MakePlay(2, 2)

		win, player := grid.HasWin()
		if *player != 2 || !win {
			t.Errorf("Win expected")
		}

		grid.Reset()
		win, player = grid.HasWin()
		if win {
			t.Errorf("Expected win to be false")
		} else if player != nil {
			t.Errorf("Expected player to be nil")
		}
	})

	t.Run("Position too large causes error", func(t *testing.T) {
		grid := tictacgo.Grid{}
		err := grid.MakePlay(9, 1)
		if err == nil {
			t.Errorf("Expected error")
		}
	})

	t.Run("Invalid player causes error", func(t *testing.T) {
		grid := tictacgo.Grid{}
		err := grid.MakePlay(0, 3)
		if err == nil {
			t.Errorf("Expected error")
		}
	})

	t.Run("Cannot play on same space", func(t *testing.T) {
		grid := tictacgo.Grid{}
		grid.MakePlay(0, 1)
		err := grid.MakePlay(0, 1)
		if err == nil {
			t.Errorf("Expected error")
		}
	})
}
