package main

import (
	"github.com/kisinga/go-htmx-tictactoe/handler"
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

	homeHandler := handler.HomeHandler{}

	app.GET("/", homeHandler.HandleHome)

	app.Static("/static", "static")

	app.Start(":8080")
}
