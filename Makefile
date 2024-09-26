run:
	go run cmd/music-lib/main.go

db-up:
	docker run --name=music-lib -e POSTGRES_PASSWORD=qwerty -p 5436:5432 -d --rm postgres

migr:
	migrate -path ./internal/db/migrations -database "postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable" up

swagger:
	swag init -g cmd/music-lib/main.go

migrate-init:
	migrate create -ext sql -dir ./internal/db/migrations -seq init