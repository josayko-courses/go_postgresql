package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID          int64
	Name, Email string
	CreatedAt   time.Time
}

type UsersModel struct {
	DB *sql.DB
}

func (m UsersModel) Insert(u *User) error {
	q := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at`
	return m.DB.QueryRow(q, u.Name, u.Email).Scan(&u.ID, &u.CreatedAt)
}
