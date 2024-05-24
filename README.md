# htmx-tictactoe

My attempt at creating a realtime online multiplayer game in golang and htmx

## How to run

I use [air](https://github.com/cosmtrek/air) to watch the golang files and live reload when I make changes


in case air does not terminate the server

```bash
sudo lsof -i :8080
```

then kill the process

```bash
kill -9 <PID>
```

## How it works

The program used htmx for the frontend and golang for the backend. The backend is a simple websocket server that listens for incoming connections. The frontend is a simple tic-tac-toe game that sends the moves to the server and the server broadcasts the moves to all the clients connected to the server.
State is maintained on the server and the server sends the state to the clients.
In order to receive state changes, the server creates a channel for each client. Each client subscribes to a specific game via SSE, meaning that each game change is broadcasted to all listening clients attached to it.
Each game is itself a channel that broadcasts the state of the game to all the clients listening to it.
An SSE channel is maintained for each client and each game. Whenever a client disconnects, the server closes the channel and removes the client from the game.

## Logic flow

Homepage --> Create Game (POST new_game, redirect --> board/gameID)
board/gameID --> Join Game (GET request to /events/gameID)  

At this point the client is listening to the gameID channel and will receive updates from the server

board/gameID --> Make Move (POST play/gameID/row/col)

Currently, any player can make a move, but the server will automatically switch the player after each move. The server will also check for a win or a draw and send the state to all the clients.

## TODO
Allow a max of 2 players to play the game, the rest should just be spectators
Fix resetting the game after a win or a draw
