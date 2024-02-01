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
	Games *map[string]*model.Board
}

func (h *PlayHandler) HandlePlay(c echo.Context) error {
	rowStr := c.QueryParam("row")
	colStr := c.QueryParam("col")
	gameID := c.QueryParam("gameID")

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
	if err != nil {
		fmt.Errorf("error: %v", err)
		return err
	}
	if winner != nil {
		c.Response().Header().Set("HX-Reload", "true")
		return c.Render(http.StatusOK, "winner", game)
	}

	return render(c, view.Cell(*cell, gameID))
}
