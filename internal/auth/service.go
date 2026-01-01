package auth

import (
	"context"
	"database/sql"
)

func CreateUser(db *sql.DB, u *User) error {
	query := `INSERT INTO users (email, name, password) VALUES (?, ?, ?)`
	result, err := db.ExecContext(context.Background(), query, u.Email, u.Name, u.Password)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.Id = int(id)

	return nil
}
