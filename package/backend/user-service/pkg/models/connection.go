package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/user-service/pkg/config"
)

const schema = `
		create table if not exists users (
			id serial primary key,
			email varchar(255) NOT NULL UNIQUE,
			password varchar(255) NOT NULL,
			created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		);
	`

var (
	db *sqlx.DB
)

func InitDatabase() error {
	DBSpec := fmt.Sprintf(
		"user=%s dbname=%s password=%s port=%s host=%s sslmode=disable",
		config.Get().DBUser, config.Get().DBName,
		config.Get().DBPassword, config.Get().DBPort,
		config.Get().DBHost,
	)
	database, err := sqlx.Connect(config.Get().DBDriver, DBSpec)
	if err != nil {
		return err
	}
	db = database
	db.MustExec(schema)
	return nil
}
