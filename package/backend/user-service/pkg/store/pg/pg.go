package pg

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

func Dial() (*sqlx.DB, error) {
	DBSpec := fmt.Sprintf(
		"user=%s dbname=%s password=%s port=%s host=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
	)
	database, err := sqlx.Connect("postgres", DBSpec)
	if err != nil {
		return &sqlx.DB{}, err
	}

	err = database.Ping()
	if err != nil {
		return &sqlx.DB{}, err
	}

	database.MustExec(schema)
	return database, nil
}
