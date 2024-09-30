# run app
run:
	go run cmd/music-lib/main.go

# docker postgres
db-up:
	docker run --name=music-lib -e POSTGRES_PASSWORD=qwerty -p 5436:5432 -d --rm postgres

swagger:
	swag init -g cmd/music-lib/main.go

# docker compose up
make up:
	docker compose -f docker-compose.yml up -d

# install linter
install-golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.3

# start linter
lint:
	golangci-lint run ./... --config .golangci.pipeline.yaml