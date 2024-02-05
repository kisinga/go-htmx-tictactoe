package handler

import (
	"github.com/kisinga/go-htmx-tictactoe/model"
	"github.com/kisinga/go-htmx-tictactoe/view"
	"github.com/labstack/echo/v4"
)

type BoardHandler struct {
	Games *map[string]*model.Board
}

func (h *BoardHandler) HandleBoard(c echo.Context) error {
	gameId := c.Param("gameID")
	game := (*h.Games)[gameId]
	if game == nil {
		return c.String(404, "game not found")
	}
	return render(c, view.Board(game))
}
