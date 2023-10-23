package main

import (
	"fmt"
	"net/http"

	"github.com/kisinga/go-htmx-tictactoe/game"
	"github.com/kisinga/go-htmx-tictactoe/template"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func main() {

	g := game.NewGame("", "")
	err := g.Play(game.X, 0, 0, game.Player1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(g)

	e := echo.New()

	// Little bit of middlewares for housekeeping
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	// This will initiate our template renderer
	template.NewTemplateRenderer(e, "public/*.html")

	e.GET("/hello", func(e echo.Context) error {
		res := map[string]interface{}{
			"Name":  "Kisinga",
			"Phone": "012345678",
			"Email": "tester@gmail.com",
		}
		return e.Render(http.StatusOK, "index", res)
	})

	e.Logger.Fatal(e.Start(":4040"))
}
