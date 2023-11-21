//go:generate yarn build:css
package main

import (
	"net/http"

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
		// gameId := e.Param("id")
		return e.Render(http.StatusOK, "index", g)
	})

	e.GET("/play", func(c echo.Context) error {
		newPlay := game.Play{}
		return g.Play(newPlay)
	})

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":4040"))
}
