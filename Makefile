generate:
	sqlc generate
	npx tailwindcss -i ./ui/css/input.css -o ./ui/static/css/styles.css 
	templ generate

build: generate
	go build main.go

run: build
	go run main.go run-api

docker-compose-up: build
	docker compose -f "docker-compose.yml" up -d --build 

docker-compose-down:
	docker compose -f "docker-compose.yml" down