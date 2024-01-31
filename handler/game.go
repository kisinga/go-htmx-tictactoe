package handler

import "github.com/labstack/echo/v4"

type GameHandler struct {
}

func (h *GameHandler) HandleGame(c echo.Context) error {
	return c.String(200, "Hello World")
}
