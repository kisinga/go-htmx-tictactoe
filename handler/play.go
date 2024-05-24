package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kisinga/go-htmx-tictactoe/model"
	"github.com/kisinga/go-htmx-tictactoe/view"
	"github.com/labstack/echo/v4"
)

type PlayHandler struct {
	Games            *map[string]*model.Board
	BroadcastChannel model.BroadcastChannel
}

func (h *PlayHandler) HandlePlay(c echo.Context) error {
	gameID := c.QueryParam("gameID")
	rowStr := c.QueryParam("row")
	colStr := c.QueryParam("col")

	row, err := strconv.Atoi(rowStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid row value")
	}

	col, err := strconv.Atoi(colStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid column value")
	}

	game := (*h.Games)[gameID]

	if game == nil {
		return c.String(http.StatusBadRequest, "invalid game id provided")
	}

	winner, cell, err := game.TakeTurn(row, col)

	// send the cell to the broadcast channel
	h.BroadcastChannel <- model.BroadcastChannelStruct{
		GameID:      gameID,
		UpdatedCell: cell,
	}

	if err != nil {
		_ = fmt.Errorf("error: %v", err)
		return err
	}
	if winner != nil {
		c.Response().Header().Set("HX-Reload", "true")
		return c.Render(http.StatusOK, "winner", game.Winner)
	}

	return render(c, view.Cell(*cell, gameID))
}
