package main

import "fmt"

type Square int

const (
	X Square = 1
	O Square = -1
)

func SquareToString(s *Square) string {
	if s == nil {
		return " "
	}
	switch *s {
	case X:
		return "X"
	case O:
		return "O"
	}
	return " "
}

type Board [][]*Square

func NewBoard() Board {
	return Board{[]*Square{nil, nil, nil}, []*Square{nil, nil, nil}, []*Square{nil, nil, nil}}
}

func (b *Board) Display() {
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
