package game

import (
	"errors"
	"fmt"
)

type xORo string

const X xORo = "x"
const O xORo = "o"

type play struct {
	owner Player
	value xORo
}

type playerNames struct {
	player1 string
	player2 string
}

type row [3]*play
type grid [3]row
type game struct {
	grid           [3]row
	nextPlayerTurn Player
	playerNames    playerNames
}

type Player = int

const (
	Player1 Player = 1 << iota
	Player2 Player = iota
)

func NewGame(player1Name, player2Name string) *game {
	return &game{
		grid: grid{
			row{
				nil, nil, nil,
			},
			row{
				nil, nil, nil,
			},
			row{
				nil, nil, nil,
			},
		},
		nextPlayerTurn: Player1,
		playerNames: playerNames{
			player1: player1Name,
			player2: player2Name,
		},
	}
}

func (g *game) Play(xoro xORo, row int, col int, p Player) error {
	if g.grid[row][col] == nil {
		g.grid[row][col] = &play{
			owner: p,
			value: xoro,
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
	for x, _ := range g.grid {
		if g.grid[0][x] != nil && g.grid[1][x] != nil && g.grid[2][x] != nil {
			if (g.grid[0][x].value == g.grid[1][x].value && g.grid[1][x].value == g.grid[2][x].value) && (g.grid[0][x].owner == g.grid[1][x].owner && g.grid[1][x].owner == g.grid[2][x].owner) {
				return &g.grid[0][x].owner
			}
		}
	}
	// 3 rows
	for x, _ := range g.grid {
		if g.grid[x][0] != nil && g.grid[x][1] != nil && g.grid[x][2] != nil {
			if (g.grid[x][0].value == g.grid[x][1].value && g.grid[x][1].value == g.grid[x][2].value) && (g.grid[x][0].owner == g.grid[x][1].owner && g.grid[x][1].owner == g.grid[x][2].owner) {
				return &g.grid[0][x].owner
			}
		}
	}

	// diagonals

	return nil
}
