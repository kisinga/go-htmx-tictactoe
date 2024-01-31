package model

import "errors"

type XORO string

const X XORO = "x"
const O XORO = "o"

type ElementValue struct {
	Owner Player
	XorO  XORO
}

type ElementID struct {
	Row int
	Col int
}
type Element struct {
	// I set this as a pointer so that I can check if it is nil or not
	Value *ElementValue
	Id    ElementID
}

type PlayerNames struct {
	Player1 string
	Player2 string
}

type Row [3]Element

type Board struct {
	GameID         int
	Grid           [3]Row
	NextPlayerTurn Player
	NextPlayerXorO XORO
	PlayerNames    PlayerNames
	Winner         *Player
}

type Player = int

const (
	Player1 Player = 1
	Player2 Player = 2
)

func (g *Board) TakeTurn(row int, col int) (winner *int, element *Element, err error) {
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return nil, nil, errors.New("illegal move")
	}
	e := Element{
		Value: &ElementValue{
			Owner: g.NextPlayerTurn,
			XorO:  g.NextPlayerXorO,
		},
		Id: ElementID{
			Row: row,
			Col: col,
		},
	}
	if g.Grid[row][col].Value == nil {
		g.Grid[row][col] = e
		if g.NextPlayerTurn == Player1 {
			g.NextPlayerTurn = Player2
		} else {
			g.NextPlayerTurn = Player1
		}
		if g.NextPlayerXorO == X {
			g.NextPlayerXorO = O
		} else {
			g.NextPlayerXorO = X
		}
	} else {
		return nil, nil, errors.New("illegal move")
	}
	winner = g.checkWinner()

	if winner != nil {
		g.Winner = winner
	}
	return winner, &e, nil
}

func (g *Board) checkWinner() *Player {
	// conditions for winning
	// 3 cols
	for x := range g.Grid {
		if g.Grid[0][x].Value != nil && g.Grid[1][x].Value != nil && g.Grid[2][x].Value != nil {
			if *g.Grid[0][x].Value == *g.Grid[1][x].Value && *g.Grid[1][x].Value == *g.Grid[2][x].Value {
				return &g.Grid[0][x].Value.Owner
			}
		}
	}
	// 3 rows
	for x := range g.Grid {
		if g.Grid[x][0].Value != nil && g.Grid[x][1].Value != nil && g.Grid[x][2].Value != nil {
			if *g.Grid[x][0].Value == *g.Grid[x][1].Value && *g.Grid[x][1].Value == *g.Grid[x][2].Value {
				return &g.Grid[0][x].Value.Owner
			}
		}
	}
	// diagonals
	if g.Grid[0][0].Value != nil && g.Grid[1][1].Value != nil && g.Grid[2][2].Value != nil {
		if *g.Grid[0][0].Value == *g.Grid[1][1].Value && *g.Grid[1][1].Value == *g.Grid[2][2].Value {
			return &g.Grid[0][0].Value.Owner
		}
	}
	if g.Grid[0][2].Value != nil && g.Grid[1][1].Value != nil && g.Grid[2][0].Value != nil {
		if *g.Grid[0][2].Value == *g.Grid[1][1].Value && *g.Grid[1][1].Value == *g.Grid[2][0].Value {
			return &g.Grid[0][2].Value.Owner
		}
	}

	return nil
}
