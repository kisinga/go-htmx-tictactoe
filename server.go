//go:generate yarn build:css
package main

import (
	"errors"
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

	g := game.NewGame("Test1", "Test2")

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

	e.GET("/game/:id", func(e echo.Context) error {
		gameId := e.Param("id")
		fmt.Println(gameId)
		err := e.Render(http.StatusOK, "index", g)
		fmt.Errorf("error: %v", err)
		return err
	})

	e.POST("new_game", func(c echo.Context) error {
		newGame := game.NewGame("Test1", "Test2")
		return c.Render(http.StatusOK, "index", newGame)
	})

	e.GET("/play", func(c echo.Context) error {
		rowStr := c.QueryParam("row")
		colStr := c.QueryParam("col")
		fmt.Println(rowStr, colStr)
		row, err := strconv.Atoi(rowStr)
		if err != nil {
			return errors.New("invalid row value")
		}
		col, err := strconv.Atoi(colStr)
		if err != nil {
			return errors.New("invalid column value")
		}
		winner, e, err := g.TakeTurn(row, col)
		if winner != nil {
			fmt.Println(winner, *e, err)
		}
		if err != nil {
			fmt.Errorf("error: %v", err)
			return err
		}
		return c.Render(http.StatusOK, "element", e)
	})

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":4040"))
}
