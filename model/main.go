package model

import "errors"

type XORO string

func (x XORO) String() string {
	return string(x)
}

type BroadcastChannelStruct struct {
	GameID      string `json:"gameID"`
	UpdatedCell *Cell  `json:"updatedCell"`
}
type BroadcastChannel chan BroadcastChannelStruct

const X XORO = "x"
const O XORO = "o"

type CellValue struct {
	Owner Player
	XorO  XORO
}

type Cell struct {
	// I set this as a pointer so that I can check if it is nil or not
	Value *CellValue
	Row   int `json:"row"`
	Col   int `json:"col"`
}

type PlayerNames struct {
	Player1 string
	Player2 string
}

type Row [3]Cell

type Board struct {
	GameID         string
	Rows           [3]Row
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

func (g *Board) TakeTurn(row int, col int) (winner *int, cell *Cell, err error) {
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return nil, nil, errors.New("illegal move")
	}
	e := Cell{
		Value: &CellValue{
			Owner: g.NextPlayerTurn,
			XorO:  g.NextPlayerXorO,
		},

		Row: row,
		Col: col,
	}
	if g.Rows[row][col].Value == nil {
		g.Rows[row][col] = e
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
	for x := range g.Rows {
		if g.Rows[0][x].Value != nil && g.Rows[1][x].Value != nil && g.Rows[2][x].Value != nil {
			if *g.Rows[0][x].Value == *g.Rows[1][x].Value && *g.Rows[1][x].Value == *g.Rows[2][x].Value {
				return &g.Rows[0][x].Value.Owner
			}
		}
	}
	// 3 rows
	for x := range g.Rows {
		if g.Rows[x][0].Value != nil && g.Rows[x][1].Value != nil && g.Rows[x][2].Value != nil {
			if *g.Rows[x][0].Value == *g.Rows[x][1].Value && *g.Rows[x][1].Value == *g.Rows[x][2].Value {
				return &g.Rows[0][x].Value.Owner
			}
		}
	}
	// diagonals
	if g.Rows[0][0].Value != nil && g.Rows[1][1].Value != nil && g.Rows[2][2].Value != nil {
		if *g.Rows[0][0].Value == *g.Rows[1][1].Value && *g.Rows[1][1].Value == *g.Rows[2][2].Value {
			return &g.Rows[0][0].Value.Owner
		}
	}
	if g.Rows[0][2].Value != nil && g.Rows[1][1].Value != nil && g.Rows[2][0].Value != nil {
		if *g.Rows[0][2].Value == *g.Rows[1][1].Value && *g.Rows[1][1].Value == *g.Rows[2][0].Value {
			return &g.Rows[0][2].Value.Owner
		}
	}

	return nil
}

func CreateNewBoard(player1Name, player2Name string, gameID string) *Board {
	g := &Board{
		GameID: gameID,
		Rows: [3]Row{
			{
				Cell{
					Value: nil,
					Row:   0, Col: 0,
				},
				Cell{
					Value: nil,
					Row:   0, Col: 1,
				},
				Cell{
					Value: nil,
					Row:   0, Col: 2,
				},
			},
			{
				Cell{
					Value: nil,
					Row:   1, Col: 0,
				},
				Cell{
					Value: nil,
					Row:   1, Col: 1,
				},
				Cell{
					Value: nil,
					Row:   1, Col: 2,
				},
			},
			{
				Cell{
					Value: nil,
					Row:   2, Col: 0,
				},
				Cell{
					Value: nil,
					Row:   2, Col: 1,
				},
				Cell{
					Value: nil,
					Row:   2, Col: 2,
				},
			},
		},
		NextPlayerTurn: Player1,
		NextPlayerXorO: X,
		PlayerNames: PlayerNames{
			Player1: player1Name,
			Player2: player2Name,
		},
	}
	return g
}
