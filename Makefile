build:
	@npm run build:css
	@templ generate
	@go run cmd/main.go

all: build