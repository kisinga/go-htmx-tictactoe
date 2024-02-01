package board

import (
	"github.com/kisinga/go-htmx-tictactoe/model"
)

func CreateNewBoard(player1Name, player2Name string, gameID string) *model.Board {
	g := &model.Board{
		GameID: gameID,
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
