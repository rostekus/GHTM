.PHONY: css run templ

DB_URL=postgresql://postgres:password@localhost:5432/mydatabase?sslmode=disable

all:
	go run cmd/main.go

css : format
	npx tailwindcss -i ./config/input.css -o ./static/css/style.css

run: templ
	go run cmd/main.go

templ: css
	templ generate

format:
	templ fmt .

postgres:
	docker run -d --name postgresql -e POSTGRES_PASSWORD=password -e POSTGRES_DB=mydatabase -p 5432:5432 postgres:latest

migrateup:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

new_migration:
	migrate create -ext sql -dir db/migrations -seq $(name)
