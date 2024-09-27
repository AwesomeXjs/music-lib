run:
	go run cmd/music-lib/main.go

db-up:
	docker run --name=music-lib -e POSTGRES_PASSWORD=qwerty -p 5436:5432 -d --rm postgres

swagger:
	swag init -g cmd/music-lib/main.go

migrate-init:
	migrate create -ext sql -dir ./internal/db/migrations -seq init

# DOCKER COMMANDS
make up:
	docker compose -f docker-compose.yml up -d

