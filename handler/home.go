package handler

import (
	"github.com/kisinga/go-htmx-tictactoe/view"
	"github.com/labstack/echo/v4"
)

// By keeping an independent handler for each route, we can easily have separetion of concerns
// and also keep the codebase maintainable and readable.
// This is also good for testing as we can easily test each handler independently.
type HomeHandler struct{}

func (h *HomeHandler) HandleHome(c echo.Context) error {
	return render(c, view.Root())
}
