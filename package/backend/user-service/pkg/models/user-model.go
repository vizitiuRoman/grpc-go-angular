package models

import (
	"time"
)

type User struct {
	ID        uint64    `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at" `
	UpdatedAt time.Time `db:"updated_at" `
}
