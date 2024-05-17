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