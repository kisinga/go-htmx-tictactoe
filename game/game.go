package game

import (
	"math/rand"

	"github.com/kisinga/go-htmx-tictactoe/model"
)

type games map[int]*model.Board

var Games games

func CreateNewGame(player1Name, player2Name string, id int) *model.Board {
	g := &model.Board{
		GameID: id,
		Grid: [3]model.Row{
			{
				model.Element{
					Value: nil,
					Id:    model.ElementID{Row: 0, Col: 0},
				},
				model.Element{
					Value: nil,
					Id:    model.ElementID{Row: 0, Col: 1},
				},
				model.Element{
					Value: nil,
					Id:    model.ElementID{Row: 0, Col: 2},
				},
			},
			{
				model.Element{
					Value: nil,
					Id:    model.ElementID{Row: 1, Col: 0},
				},
				model.Element{
					Value: nil,
					Id:    model.ElementID{Row: 1, Col: 1},
				},
				model.Element{
					Value: nil,
					Id:    model.ElementID{Row: 1, Col: 2},
				},
			},
			{
				model.Element{
					Value: nil,
					Id:    model.ElementID{Row: 2, Col: 0},
				},
				model.Element{
					Value: nil,
					Id:    model.ElementID{Row: 2, Col: 1},
				},
				model.Element{
					Value: nil,
					Id:    model.ElementID{Row: 2, Col: 2},
				},
			},
		},
		NextPlayerTurn: model.Player1,
		NextPlayerXorO: model.X,
		PlayerNames: model.PlayerNames{
			Player1: player1Name,
			Player2: player2Name,
		},
	}
	return g
}

func NewGame(player1Name, player2Name string) int {
	id := generateGameId()
	g := CreateNewGame(player1Name, player2Name, id)
	Games[id] = g
	return id
}

// a function that generates a random game id
func generateGameId() int {
	return rand.Intn(100000)
}

func init() {
	Games = make(games)
}
