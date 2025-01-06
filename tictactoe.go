package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

type Square int

const (
	X     Square = 1
	O     Square = -1
	BLANK Square = 0
)

func SquareToString(s Square) string {
	switch s {
	case X:
		return "X"
	case O:
		return "O"
	default:
		return " "
	}
}

type Board [][]Square

func NewBoard(rows, cols int) *Board {
	board := Board{}
	for i := 0; i < rows; i++ {
		row := []Square{}
		for j := 0; j < cols; j++ {
			row = append(row, BLANK)
		}
		board = append(board, row)
	}

	return &board
}

// TODO: This assumes 3x3 board
func (b *Board) ToString() string {
	if b == nil {
		return "BOARD IS NIL"
	}

	s := "_____________\n"
	for i, row := range *b {
		if i != 0 {
			s += "+———+———+———+\n"
		}
		s += "|"
		for _, square := range row {
			s += fmt.Sprintf(" %s |", SquareToString(square))
		}
		s += "\n"
	}
	s += "‾‾‾‾‾‾‾‾‾‾‾‾\n"
	return s
}

func (b *Board) IsLegal(m Move) error {
	if m.Mark != X && m.Mark != O {
		return errors.New("mark must be either X or O")
	}

	if int(m.Row) > len(*b) || int(m.Col) > len((*b)[m.Row]) {
		return fmt.Errorf(
			"row %d Column %d is out of bounds for board with dimensions %dx%d", len(*b), len((*b)[0]), m.Row, m.Col,
		)
	}

	if (*b)[m.Row][m.Col] != BLANK {
		return fmt.Errorf(
			"row %d Column %d already has a %s", m.Row, m.Col, SquareToString((*b)[m.Row][m.Col]),
		)
	}

	return nil
}

func (b *Board) Play(move Move) error {
	if b == nil {
		return errors.New("Board is nil")
	}

	if err := b.IsLegal(move); err != nil {
		return err
	}

	(*b)[move.Row][move.Col] = move.Mark
	return nil
}

// Assumes 3x3 board
func (b *Board) CheckWinner() Square {
	if b == nil {
		return BLANK
	}

	sums := b.computeSums()
	for _, sum := range sums {
		if sum == 3 {
			return X
		}

		if sum == -3 {
			return O
		}
	}

	return BLANK
}

func (b Board) computeSums() []int {
	r0 := int(b[0][0] + b[0][1] + b[0][2])
	r1 := int(b[1][0] + b[1][1] + b[1][2])
	r2 := int(b[2][0] + b[2][1] + b[2][2])

	c0 := int(b[0][0] + b[1][0] + b[2][0])
	c1 := int(b[0][1] + b[1][1] + b[2][1])
	c2 := int(b[0][2] + b[1][2] + b[2][2])

	d0 := int(b[0][0] + b[1][1] + b[2][2])
	d1 := int(b[0][2] + b[1][1] + b[2][0])

	return []int{r0, r1, r2, c0, c1, c2, d0, d1}
}

func (b Board) isFull() bool {
	for _, row := range b {
		for _, square := range row {
			if square == BLANK {
				return false
			}
		}
	}
	return true
}

func (b *Board) clone() *Board {
	clone := NewBoard(len(*b), len((*b)[0]))
	for rowNum, row := range *b {
		for colNum := range row {
			(*clone)[rowNum][colNum] = (*b)[rowNum][colNum]
		}
	}

	return clone
}

func (b *Board) listLegalMoves(objective Square) []Move {
	moves := []Move{}

	// Find all legal moves
	for rowNum, row := range *b {
		for col := range row {
			if row[col] == BLANK {
				moves = append(moves, Move{rowNum, col, objective})
			}
		}
	}

	return moves
}

type Move struct {
	Row  int
	Col  int
	Mark Square
}

type Strategy func(b *Board, objective Square) Move

func NewRandomStrategy() Strategy {
	return func(b *Board, objective Square) Move {
		for {
			move := Move{rand.Intn(len(*b)), rand.Intn(len((*b)[0])), objective}
			if err := b.IsLegal(move); err == nil {
				return move
			}
		}
	}
}

func NewInformedStrategy() Strategy {
	return func(b *Board, objective Square) Move {
		candidateMoves := b.listLegalMoves(objective)

		var bestMove Move
		bestStrength := -math.MaxInt

		// Check all moves and find the one which wins or maximizes the row, column, and diagonal scores
		for _, move := range candidateMoves {
			boardCopy := b.clone()
			boardCopy.Play(move)
			if boardCopy.CheckWinner() == objective {
				return move
			}

			sums := boardCopy.computeSums()
			strength := 0
			for _, sum := range sums {
				// multiply by 1 if X, -1 if O
				strength += sum * int(objective)
			}

			if strength > bestStrength {
				bestMove = move
				bestStrength = strength
			}
		}

		return bestMove
	}
}

func NewSmartStrategy() Strategy {
	return func(b *Board, objective Square) Move {
		candidateMoves := b.listLegalMoves(objective)

		var bestMove *Move = nil
		myBestScore := -math.MaxInt
		opponentBestScore := -math.MaxInt

		for _, move := range candidateMoves {
			boardCopy := b.clone()
			boardCopy.Play(move)
			if boardCopy.CheckWinner() == objective {
				return move
			}

			myBestSum := -math.MaxInt
			opponentBestSum := math.MaxInt

			sums := boardCopy.computeSums()
			for _, sum := range sums {
				// weightedSum is the row, column, or diagonal sum multiplied by 1 for X or -1 for O
				weightedSum := sum * int(objective)
				if weightedSum > myBestSum {
					myBestSum = weightedSum
				}

				if weightedSum < opponentBestSum {
					opponentBestSum = weightedSum
				}
			}

			if myBestSum == 3 {
				return move
			}

			if opponentBestSum >= opponentBestScore {
				bestMove = &move
				opponentBestScore = opponentBestSum
			}

			if myBestSum >= myBestScore && opponentBestSum >= opponentBestScore {
				bestMove = &move
				myBestScore = myBestSum
				opponentBestScore = opponentBestSum
			}

		}

		centerMove := Move{1, 1, objective}
		if b.IsLegal(centerMove) == nil {
			boardCopy := b.clone()
			boardCopy.Play(centerMove)
			myBestSum := -math.MaxInt
			opponentBestSum := math.MaxInt
			sums := boardCopy.computeSums()
			for _, sum := range sums {
				// weightedSum is the row, column, or diagonal sum multiplied by 1 for X or -1 for O
				weightedSum := sum * int(objective)
				if weightedSum > myBestSum {
					myBestSum = weightedSum
				}

				if weightedSum < opponentBestSum {
					opponentBestSum = weightedSum
				}
			}

			if myBestSum >= myBestScore && opponentBestSum >= opponentBestScore {
				bestMove = &centerMove
			}
		}

		return *bestMove
	}
}

type TicTacToePlayer struct {
	Mark     Square
	Strategy Strategy
}

func NewTicTacToePlayer(mark Square, strat Strategy) *TicTacToePlayer {
	if mark != X && mark != O {
		return nil
	}

	return &TicTacToePlayer{mark, strat}
}

func (t *TicTacToePlayer) ChooseMove(b *Board) Move {
	return t.Strategy(b, t.Mark)
}

type TicTacToeGame struct {
	PlayerX *TicTacToePlayer
	PlayerO *TicTacToePlayer
	Board   *Board
}

func NewTicTacToeGame(playerX, playerO *TicTacToePlayer) TicTacToeGame {
	return TicTacToeGame{playerX, playerO, NewBoard(3, 3)}
}

func (g *TicTacToeGame) Play() *TicTacToePlayer {
	for {
		moveX := g.PlayerX.ChooseMove(g.Board)

		if err := g.Board.Play(moveX); err != nil {
			fmt.Printf("X made an illegal move: %s\n", err)
			return g.PlayerO
		}

		fmt.Println("X's move:")
		fmt.Print(g.Board.ToString())

		winner := g.Board.CheckWinner()

		if winner == X {
			return g.PlayerX
		}

		if g.Board.isFull() && winner == BLANK {
			return nil
		}

		moveO := g.PlayerO.ChooseMove(g.Board)

		if err := g.Board.Play(moveO); err != nil {
			fmt.Printf("O made an illegal move: %s\n", err)
			return g.PlayerX
		}

		fmt.Println("O's move:")
		fmt.Print(g.Board.ToString())

		winner = g.Board.CheckWinner()

		if winner == O {
			return g.PlayerO
		}

		if g.Board.isFull() && winner == BLANK {
			return nil
		}
	}
}
