generate:
	sqlc generate
	templ generate

build: generate
	go build main.go


run: build
	go run main.go run-api