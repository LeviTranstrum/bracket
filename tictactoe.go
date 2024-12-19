package main

import (
	"errors"
	"fmt"
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

func NewBoard(rows, cols uint) *Board {
	board := Board{}
	for i := uint(0); i < rows; i++ {
		row := []Square{}
		for j := uint(0); j < cols; j++ {
			row = append(row, BLANK)
		}
		board = append(board, row)
	}

	return &board
}

// TODO: This assumes 3x3 board
func (b *Board) Display() {
	if b == nil {
		fmt.Println("BOARD IS NIL")
		return
	}

	fmt.Println("_____________")
	for i, row := range *b {
		if i != 0 {
			fmt.Println("+———+———+———+")
		}
		fmt.Print("|")
		for _, square := range row {
			fmt.Printf(" %s |", SquareToString(square))
		}
		fmt.Print("\n")
	}
	fmt.Println("‾‾‾‾‾‾‾‾‾‾‾‾‾")
}

func (b *Board) IsLegal(m Move) error {
	if m.Mark != X && m.Mark != O {
		return errors.New("Mark must be either X or O")
	}

	if int(m.Row) > len(*b) || int(m.Col) > len((*b)[m.Row]) {
		return errors.New(fmt.Sprintf(
			"Row %d Column %d is out of bounds for board with dimensions %dx%d", len(*b), len((*b)[0]),
		))
	}

	if (*b)[m.Row][m.Col] != BLANK {
		return errors.New(fmt.Sprintf(
			"Row %d Column %d already has a %s", m.Row, m.Col, SquareToString((*b)[m.Row][m.Col]),
		))
	}

	return nil
}

// TODO: Pull some of this out into b.IsLegal(move)
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
func (b Board) CheckWinner() *Square {
	if b == nil {
		return nil
	}

	sums := b.computeSums()
	for _, sum := range sums {
		if sum == 3 {
			x := X
			return &x
		}

		if sum == -3 {
			o := O
			return &o
		}
	}

	if b.isFull() {
		blank := BLANK
		return &blank
	}

	return nil
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

type Move struct {
	Row  uint
	Col  uint
	Mark Square
}

type Strategy func(b *Board, objective Square) Move

func NewRandomStrategy() Strategy {
	return func(b *Board, objective Square) Move {
		for {
			move := Move{uint(rand.Intn(len(*b))), uint(rand.Intn(len((*b)[0]))), objective}
			if err := b.IsLegal(move); err == nil {
				return move
			}
		}
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
		g.Board.Display()

		if winner := g.Board.CheckWinner(); winner != nil && *winner == X {
			return g.PlayerX
		}

		moveO := g.PlayerO.ChooseMove(g.Board)

		if err := g.Board.Play(moveO); err != nil {
			fmt.Printf("O made an illegal move: %s\n", err)
			return g.PlayerX
		}

		fmt.Println("O's move:")
		g.Board.Display()

		winner := g.Board.CheckWinner()

		if winner != nil && *winner == O {
			return g.PlayerO
		}

		if winner != nil && *winner == BLANK {
			return nil
		}
	}
}
