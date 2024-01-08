package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"login-ports/lib/env"

	"github.com/pressly/goose"
)

func Migrate() {
	var db *sql.DB

	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		env.String("Postgres.Host", "localhost"),
		env.String("Postgres.Port", "5432"),
		env.String("Postgres.Database", "login-ports"),
		env.String("Postgres.User", "user"),
		env.String("Postgres.Password", "securePassword"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Println(err)
		return
	}

	if err := goose.Run("up", db, "migrations"); err != nil {
		log.Println(err)
		return
	}
}
