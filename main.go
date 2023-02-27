package main

import (
	"fmt"
)

type gameState struct {
	Board  [3][3]int
	Player int
}

const (
	none     = iota
	cross    = iota
	circle   = iota
	noWinner = iota
	winnerX
	winnerO
	draw
)

func main() {
	state := gameState{}
	state.Player = cross
	// result := noWinner

	for {
		fmt.Printf("next player to place a mark is: %v\n", printPlayer(state.whosNext()))
		state.drawBoard()
		fmt.Printf("where to place a %v? (input row then column, separated by space)\n> ", printPlayer(state.whosNext()))

		for {
			var row, column int
			fmt.Scan(&row, &column)

			err := state.updateBoard(row-1, column-1) // -1 so coordinate starts at (1,1) instead of (0,0)

			// if valid position was entered, break out from the input loop
			if err == nil {
				break
			}

			// if an invalid position was entered, prompt the player to re-enter another position
			fmt.Println(err)
			fmt.Printf("please re-enter a position:\n> ")
		}

		// todo check if anyone has won the game

		// 4. if no one has won in this turn, go on for next turn and continue the game loop
		state.nextTurn()

		fmt.Println()

		// remove when implement winner check

		state.drawBoard()
	}

	// switch result {
	// case winnerX:
	// 	fmt.Printf("X won the game!\n")
	// case winnerO:
	// 	fmt.Printf("O won the game!\n")
	// case draw:
	// 	fmt.Printf("the game has ended with a draw!\n")
	// }
}

func printPlayer(player int) string {
	switch player {
	case none:
		return "none"
	case cross:
		return "X"
	case circle:
		return "O"
	default:
		return fmt.Sprintf("%d", int(player))
	}
}

// func clearScreen() {
// 	c := exec.Command("cmd", "/c", "cls")
// 	c.Stdout = os.Stdout
// 	c.Run()
// }

func (s *gameState) drawBoard() {

	for i, row := range s.Board {
		for j, column := range row {
			fmt.Print(" ")
			switch column {
			case none:
				fmt.Print(" ")
			case cross:
				fmt.Print("X")
			case circle:
				fmt.Print("O")
			}
			if j != len(row)-1 {
				fmt.Print(" |")
			}
		}
		if i != len(s.Board)-1 {
			fmt.Print("\n------------")
		}
		fmt.Print("\n")
	}
}

func (state *gameState) whosNext() int {
	return state.Player
}

func (s *gameState) nextTurn() {
	if s.Player == cross {
		s.Player = circle
	} else {
		s.Player = cross
	}
}

func (s *gameState) updateBoard(row int, column int) error {
	if row < 0 || column < 0 || row >= len(s.Board) || column >= len(s.Board[row]) {
		return fmt.Errorf("position (%d,%d) is out of bound.", row, column)
	}
	if s.Board[row][column] != none {
		return fmt.Errorf("position (%d,%d) has already been played.", row, column)
	}
	s.Board[row][column] = int(s.Player)
	return nil
}
