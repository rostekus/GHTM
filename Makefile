.PHONY: css run templ
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

mariadbup:
	docker run -d --name mariadb -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=mydatabase -p 3306:3306 mariadb:latest

migrateup:
	migrate -path db/migrations -database "mysql://root:password@tcp(127.0.0.1:3306)/mydatabase" -verbose up

migratedown:
	migrate -path db/migrations -database "mysql://root:password@tcp(127.0.0.1:3306)/mydatabase" -verbose down
