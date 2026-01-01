package auth

import (
	"context"
	"errors"
)

type UserServiceCreator interface {
	CreateUser(ctx context.Context, email string, name string, password string) error
	AuthenticateUser(ctx context.Context, email string, password string) (*LoginResponse, error)
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

	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	defer txStore.Rollback()
	user := User{Email: email, Name: name, Password: hashedPassword}
	if err := txStore.Save(ctx, user); err != nil {
		return err
	}

	if err := txStore.Commit(); err != nil {
		return err
	}

	return nil
}

type LoginResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

func (s *UserService) AuthenticateUser(ctx context.Context, email string, password string) (*LoginResponse, error) {
	txStore, err := s.store.BeginTx(ctx)
	if err != nil {
		return nil, err
	}

	defer txStore.Rollback()
	user, err := txStore.FindByEmail(ctx, email)
	if err := txStore.Commit(); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	token, err := GenerateJWT(user.Id, user.Email)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{User: *user, Token: token}, nil
}
