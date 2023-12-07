//go:generate yarn build:css
package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kisinga/go-htmx-tictactoe/game"
	"github.com/kisinga/go-htmx-tictactoe/template"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func main() {

	e := echo.New()

	// Little bit of middlewares for housekeeping
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	// This will initiate our template renderer
	template.NewTemplateRenderer(e, "public/*.html")

	e.GET("/", func(e echo.Context) error {
		return e.Render(http.StatusOK, "index", nil)
	})

	e.GET("/game/:id", func(c echo.Context) error {

		gameId := c.Param("id")
		id, err := strconv.Atoi(gameId)
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/")
		}
		g, ok := game.Games[id]
		if !ok {
			return c.Redirect(http.StatusSeeOther, "/")
		}
		err = c.Render(http.StatusOK, "index", g)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return nil
	})

	e.POST("new_game", func(c echo.Context) error {
		player1 := c.FormValue("player1")
		player2 := c.FormValue("player2")
		gameID := game.NewGame(player1, player2)
		fmt.Println(gameID)
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/game/%d", gameID))
	})

	e.GET("/play", func(c echo.Context) error {
		rowStr := c.QueryParam("row")
		colStr := c.QueryParam("col")
		gameIDStr := c.QueryParam("gameID")
		row, err := strconv.Atoi(rowStr)
		if err != nil {
			return c.String(http.StatusBadRequest, "invalid row value")
		}

		col, err := strconv.Atoi(colStr)
		if err != nil {
			return c.String(http.StatusBadRequest, "invalid column value")
		}

		gameID, err := strconv.Atoi(gameIDStr)
		if err != nil {
			return c.String(http.StatusBadRequest, "invalid game id")
		}

		g, ok := game.Games[gameID]
		if !ok {
			return c.String(http.StatusBadRequest, "invalid game id")
		}
		winner, e, err := g.TakeTurn(row, col)
		if err != nil {
			fmt.Errorf("error: %v", err)
			return err
		}
		if winner != nil {
			fmt.Println(winner, *e, err)
		}

		return c.Render(http.StatusOK, "element", e)
	})

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":4040"))
}
