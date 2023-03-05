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
	result := noWinner

	for {
		fmt.Printf("Current player: %v\n", printPlayer(state.currentPlayer()))
		state.drawBoard()
		fmt.Printf("How to place a %v? (input row then column, separated by space)\n> ", printPlayer(state.currentPlayer()))

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

		// check for winner
		result = state.computeWinner()
		if result != noWinner {
			break
		}

		// if no one has won move to the next player and continue the game loop
		state.nextTurn()

		fmt.Println()
	}

	state.drawBoard()
	switch result {
	case winnerX:
		fmt.Printf("X won the game!\n")
	case winnerO:
		fmt.Printf("O won the game!\n")
	case draw:
		fmt.Printf("DRAW!\n")
	}
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

func (state *gameState) currentPlayer() int {
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
		return fmt.Errorf("position (%d,%d) is out of bounds.", row, column)
	}
	if s.Board[row][column] != none {
		return fmt.Errorf("position (%d,%d) has already been played.", row, column)
	}
	s.Board[row][column] = int(s.Player)
	return nil
}

func (s *gameState) validateSection(l int, startingRow int, startingColumn int, deltaRow int, deltaColumn int) int {
	var lastSquare int = s.Board[startingRow][startingColumn]
	row, column := startingRow+deltaRow, startingColumn+deltaColumn

	for row >= 0 && column >= 0 && row < l && column < l {
		if s.Board[row][column] == none {
			return noWinner
		}

		if lastSquare != s.Board[row][column] {
			return noWinner
		}

		lastSquare = s.Board[row][column]
		row, column = row+deltaRow, column+deltaColumn
	}

	if lastSquare == cross {
		return winnerX
	} else if lastSquare == circle {
		return winnerO
	}
	return noWinner
}

func (s *gameState) computeWinner() int {
	boardLen := len(s.Board)

	// check horizontals rows
	for row := 0; row < boardLen; row++ {
		if result := s.validateSection(boardLen, row, 0, 0, 1); result != noWinner {
			return result
		}
	}

	// check vertical columns
	for column := 0; column < boardLen; column++ {
		if result := s.validateSection(boardLen, 0, column, 1, 0); result != noWinner {
			return result
		}
	}

	// draw case
	for _, row := range s.Board {
		for _, square := range row {
			if square == 0 {
				return noWinner
			}
		}
	}

	return draw
}
