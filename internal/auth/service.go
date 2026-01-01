package auth

import (
	"context"
)

type UserServiceCreator interface {
	CreateUser(ctx context.Context, email string, name string, password string) error
}

type UserService struct {
	store Store
}

func NewUserService(store Store) *UserService {
	return &UserService{store}
}

func (s *UserService) CreateUser(ctx context.Context, email string, name string, password string) error {
	txStore, err := s.store.BeginTx(ctx)
	if err != nil {
		return err
	}

	defer txStore.Rollback()
	user := User{Email: email, Name: name, Password: password}
	if err := txStore.Save(ctx, user); err != nil {
		return err
	}

	if err := txStore.Commit(); err != nil {
		return err
	}

	return nil
}
