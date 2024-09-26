package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

func main() {
	m, err := migrate.New(
		"github.com/AwesomeXjs/music-lib/internal/db/migrations",
		"postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	m.Up()
}
