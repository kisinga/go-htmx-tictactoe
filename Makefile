run:
	@npm run build:css
	@templ generate
	@go run cmd/main.go