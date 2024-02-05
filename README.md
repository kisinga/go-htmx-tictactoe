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
