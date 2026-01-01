package auth

import "context"

type Store interface {
	BeginTx(ctx context.Context) (TxStore, error)
}

type TxStore interface {
	Save(ctx context.Context, user User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	Commit() error
	Rollback() error
}
