package game

import (
	"errors"
	"fmt"
)

type XORO string

const X XORO = "x"
const O XORO = "o"

type Element struct {
	Owner Player
	Value XORO
}

type PlayerNames struct {
	Player1 string
	Player2 string
}

type Row [3]*Element
type game struct {
	Grid           [3]Row
	NextPlayerTurn Player
	PlayerNames    PlayerNames
}

type Player = int

const (
	Player1 Player = 1 << iota
	Player2 Player = iota
)

func NewGame(player1Name, player2Name string) *game {
	return &game{
		Grid: [3]Row{
			{
				nil, nil, nil,
			},
			{
				nil, nil, nil,
			},
			{
				nil, nil, nil,
			},
		},
		NextPlayerTurn: Player1,
		PlayerNames: PlayerNames{
			Player1: player1Name,
			Player2: player2Name,
		},
	}
}

type Play struct {
	Owner Player
	Value XORO
	row   int
	col   int
}

func (g *game) Play(p Play) error {
	if g.Grid[p.row][p.col] == nil {
		g.Grid[p.row][p.col] = &Element{
			Owner: p.Owner,
			Value: p.Value,
		}
	} else {
		return errors.New("illegal move")
	}

	winner := g.checkWinner()
	fmt.Println(winner)
	return nil
}

func (g *game) checkWinner() *Player {
	// conditions for winning
	// 3 cols
	for x, _ := range g.Grid {
		if g.Grid[0][x] != nil && g.Grid[1][x] != nil && g.Grid[2][x] != nil {
			if (g.Grid[0][x].Value == g.Grid[1][x].Value && g.Grid[1][x].Value == g.Grid[2][x].Value) && (g.Grid[0][x].Owner == g.Grid[1][x].Owner && g.Grid[1][x].Owner == g.Grid[2][x].Owner) {
				return &g.Grid[0][x].Owner
			}
		}
	}
	// 3 rows
	for x, _ := range g.Grid {
		if g.Grid[x][0] != nil && g.Grid[x][1] != nil && g.Grid[x][2] != nil {
			if (g.Grid[x][0].Value == g.Grid[x][1].Value && g.Grid[x][1].Value == g.Grid[x][2].Value) && (g.Grid[x][0].Owner == g.Grid[x][1].Owner && g.Grid[x][1].Owner == g.Grid[x][2].Owner) {
				return &g.Grid[0][x].Owner
			}
		}
	}

	// diagonals

	return nil
}
