package handler

import (
	"bytes"
	"fmt"
	"time"

	"github.com/kisinga/go-htmx-tictactoe/model"
	"github.com/kisinga/go-htmx-tictactoe/view"
	"github.com/labstack/echo/v4"
)

type EventsHandler struct {
	Games            *map[string]*model.Board
	BroadcastChannel model.BroadcastChannel
	Clients          map[string][]chan *model.Cell // Maps game IDs to client channels
}

func NewEventsHandler(games *map[string]*model.Board, bc model.BroadcastChannel, clients map[string][]chan *model.Cell) *EventsHandler {
	return &EventsHandler{
		Games:            games,
		BroadcastChannel: bc,
		Clients:          clients,
	}
}

// ListenToBroadcasts starts listening to the broadcast channel and forwards events to relevant clients.
func (h *EventsHandler) ListenToBroadcasts() {
	go func() {
		for update := range h.BroadcastChannel {
			if clients, ok := h.Clients[update.GameID]; ok {
				for _, client := range clients {
					// Non-blocking send with select to avoid blocking if client is not listening
					select {
					case client <- update.UpdatedCell:
					default:
					}
				}
			}
		}
	}()
}

// HandleEvents sets up an SSE connection and listens for updates to send to the client.
func (h *EventsHandler) HandleEvents(c echo.Context) error {
	gameID := c.Param("gameID")

	// Set up the SSE headers
	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set(echo.HeaderCacheControl, "no-cache")
	c.Response().Header().Set(echo.HeaderConnection, "keep-alive")
	c.Response().Flush()

	// Create a channel to receive updates for this client
	clientChan := make(chan *model.Cell)
	h.Clients[gameID] = append(h.Clients[gameID], clientChan)

	// Keep-alive ticker
	ticker := time.NewTicker(30 * time.Second) // Send a keep-alive every 30 seconds

	// Remove client channel when done
	defer func() {
		ticker.Stop()
		// Logic to safely remove the clientChan from h.Clients[gameID]
		for i, ch := range h.Clients[gameID] {
			if ch == clientChan {
				h.Clients[gameID] = append(h.Clients[gameID][:i], h.Clients[gameID][i+1:]...)
				break
			}
		}
		close(clientChan) // Close the channel to prevent leaks
	}()

	// Send initial connection confirmation to the client
	if _, err := fmt.Fprintf(c.Response(), "data: %s\n\n", `"connected"`); err != nil {
		return nil // Stop if the initial message fails
	}
	c.Response().Flush()

	// Listen for messages on the client channel and send them as SSE
	for {
		select {
		case cell := <-clientChan:
			if cell == nil { // Channel was closed
				return nil
			}
			var buf bytes.Buffer
			err := view.Cell(*cell, gameID).Render(c.Request().Context(), &buf)
			if err != nil {
				c.Logger().Error("Error rendering cell: ", err)
				return nil // Exit if there's an error rendering the cell
			}
			if _, err := fmt.Fprintf(c.Response(), "event:%d-%d\ndata: %v\n\n", cell.Row, cell.Col, buf.String()); err != nil {
				c.Logger().Error("Error sending SSE: ", err)
				return nil // Exit if there's an error sending the message
			}
			c.Response().Flush()
		case <-ticker.C:
			// Send a keep-alive message
			if _, err := fmt.Fprintf(c.Response(), "data: %s\n\n", `"keep-alive"`); err != nil {
				c.Logger().Error("Error sending keep-alive: ", err)
				return nil // Exit if there's an error sending the keep-alive message
			}
			c.Response().Flush()
		}
	}
}
