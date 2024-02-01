build:
	@npm run build:css
	@templ generate
	@go build -o ./tictactoe ./cmd

run: build
	@./tictactoe

all: build
