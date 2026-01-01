package auth

import (
	"context"
	"database/sql"
)

type MySqlStore struct {
	db *sql.DB
}

type MySqlTxStore struct {
	tx *sql.Tx
}

func NewMySqlStore(db *sql.DB) *MySqlStore {
	return &MySqlStore{
		db,
	}
}

func (s *MySqlStore) BeginTx(ctx context.Context) (TxStore, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &MySqlTxStore{
		tx,
	}, nil
}

func (s *MySqlTxStore) Save(ctx context.Context, u User) error {
	query := `INSERT INTO users (email, name, password) VALUES (?, ?, ?)`
	_, err := s.tx.ExecContext(ctx, query, u.Email, u.Name, u.Password)
	if err != nil {
		return err
	}

	return nil
}

func (s *MySqlTxStore) Commit() error {
	return s.tx.Commit()
}

func (s *MySqlTxStore) Rollback() error {
	return s.tx.Rollback()
}
