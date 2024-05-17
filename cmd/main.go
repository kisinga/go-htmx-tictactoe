package main

import (
	"log"

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
	broadcastChannel := make(model.BroadcastChannel)
	clients := make(map[string][]chan *model.Cell)

	games["test"] = model.CreateNewBoard("player1", "player2", "test")

	homeHandler := handler.HomeHandler{}

	playHandler := handler.PlayHandler{
		Games:            &games,
		BroadcastChannel: broadcastChannel,
	}

	newGameHandler := handler.NewGameHandler{
		Games: &games,
	}

	boardHandler := handler.BoardHandler{
		Games: &games,
	}

	eventsHandler := handler.NewEventsHandler(&games, broadcastChannel, clients)
	eventsHandler.ListenToBroadcasts()

	app.GET("/", homeHandler.HandleHome)

	app.POST("/play", playHandler.HandlePlay)

	app.POST("new_game", newGameHandler.HandleNewGame)

	app.GET("/board/:gameID", boardHandler.HandleBoard)

	// @TODO: finish resetting the game
	app.POST("/reset/:gameID", boardHandler.HandleBoard)

	app.GET("/events/:gameID", eventsHandler.HandleEvents)

	app.Static("/static", "static")

	err := app.Start(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
