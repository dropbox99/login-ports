package config

import (
	"fmt"
	"log"
	"login-ports/lib/env"
	"os"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Postgresql() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // Disable color
			},
		),
	})
}

func dsn() string {
	host := "host=" + env.String("Postgres.Host", "localhost")
	port := "port=" + env.String("Postgres.Port", "5432")
	dbname := "dbname=" + env.String("Postgres.Database", "login-ports")
	user := "user=" + env.String("Postgres.User", "user")
	password := "password=" + env.String("Postgres.Password", "securePassword")
	return fmt.Sprintln(host, port, dbname, user, password)
}
