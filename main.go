package main

import (
	"fmt"
	"math/rand"
)

type Player interface {
	TakeTurn()
}

type Game interface {
	NewGame()
	Play()
}

type GameMaker interface {
	MakeGame(...*Player)
}

type Bracket struct {
	Players   []Player
	GameMaker GameMaker
}

type BaseballPlayer struct {
	Name  string
	Skill float64
	Luck  float64
}

func NewBaseballPlayer(name string, skill, luck float64) *BaseballPlayer {
	return &BaseballPlayer{name, skill, luck}
}

func (b *BaseballPlayer) TakeTurn() {
}

type BaseballGame struct {
	P1 *BaseballPlayer
	P2 *BaseballPlayer
}

func NewBaseballGame(p1, p2 *BaseballPlayer) BaseballGame {
	return BaseballGame{p1, p2}
}

func (g *BaseballGame) Play() *BaseballPlayer {
	fmt.Printf("The %s play the %s \n", g.P1.Name, g.P2.Name)
	if g.P1.Skill+g.P1.Luck*rand.Float64() > g.P2.Skill+g.P2.Luck*rand.Float64() {
		fmt.Printf("The %s win \n", g.P1.Name)
		return g.P1
	}

	fmt.Printf("The %s win \n", g.P2.Name)
	return g.P2
}

type BaseballBracket []*BaseballPlayer

func (b BaseballBracket) PlayBracket() *BaseballPlayer {
	bracketSize := len(b)
	if bracketSize == 0 {
		return nil
	}

	nextBracket := BaseballBracket{}
	for i := 0; i < bracketSize; i += 2 {
		if i+i < bracketSize {
			game := NewBaseballGame(b[i], b[i+1])
			nextBracket = append(nextBracket, game.Play())
		} else {
			nextBracket = append(BaseballBracket{b[i]}, nextBracket...)
		}
	}

	if len(nextBracket) == 1 {
		return nextBracket[0]
	}

	return nextBracket.PlayBracket()
}

func main() {
	// RedSox := NewBaseballPlayer("Sox", 100, 0)
	// Yankees := NewBaseballPlayer("Yankees", 90, 20)
	// Angels := NewBaseballPlayer("Angels", 80, 40)
	// Dodgers := NewBaseballPlayer("Dodgers", 70, 60)
	// Lakers := NewBaseballPlayer("Lakers", 60, 100)

	// bracket := BaseballBracket{}
	// bracket = append(bracket, RedSox, Yankees, Angels, Dodgers, Lakers)

	// winner := bracket.PlayBracket()
	// fmt.Printf("The %s take the bracket\n", winner.Name)
	xWins := 0
	oWins := 0
	draws := 0
	for range 10000 {
		px := NewTicTacToePlayer(X, NewSmartStrategy())
		pO := NewTicTacToePlayer(O, NewRandomStrategy())
		game := NewTicTacToeGame(px, pO)
		winner := game.Play()
		if winner == game.PlayerX {
			fmt.Println("X wins!")
			xWins++
		}
		if winner == game.PlayerO {
			fmt.Println("O wins!")
			oWins++
		}
		if winner == nil {
			fmt.Println("Cat's game!")
			draws++
		}
	}
	fmt.Printf("xWins: %d", xWins)
	fmt.Printf("oWins: %d", oWins)
	fmt.Printf("draws: %d", draws)
}
