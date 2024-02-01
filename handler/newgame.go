package handler

import (
	"fmt"
	"net/http"

	"github.com/kisinga/go-htmx-tictactoe/board"
	"github.com/kisinga/go-htmx-tictactoe/model"
	"github.com/labstack/echo/v4"
)

type NewGameHandler struct {
	Games *map[string]*model.Board
}

func (h *NewGameHandler) HandleNewGame(c echo.Context) error {
	player1 := c.FormValue("player1")
	player2 := c.FormValue("player2")
	game := board.CreateNewBoard(player1, player2, "test")

	// add the game to the map
	(*h.Games)["game.GameID"] = game

	// fmt.Println(gameID)
	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/game/%s", game.GameID))
}
