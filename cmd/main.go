package main

import (
	"github.com/kisinga/go-htmx-tictactoe/handler"
	"github.com/kisinga/go-htmx-tictactoe/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func main() {

	app := echo.New()

	// Little bit of middlewares for housekeeping
	app.Pre(middleware.RemoveTrailingSlash())
	app.Use(middleware.Recover())
	app.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	// games state
	games := make(map[string]*model.Board)

	homeHandler := handler.HomeHandler{}

	playHandler := handler.PlayHandler{
		Games: &games,
	}

	newGameHandler := handler.NewGameHandler{
		Games: &games,
	}

	gameHandler := handler.GameHandler{}

	app.GET("/", homeHandler.HandleHome)

	app.GET("/play/:id", playHandler.HandlePlay)

	app.POST("new_game", newGameHandler.HandleNewGame)

	app.Static("/static", "static")

	app.Start(":8080")
}
