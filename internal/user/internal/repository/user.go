package repository

import (
	"context"

	"github.com/achmad-dev/internal/user/internal/domain"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetUser(ctx context.Context, id string) (*domain.User, error)
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id string) error
}

type userRepositoryImpl struct {
	sqldb *sqlx.DB
}

// GetUserByUsername implements UserRepository.
func (u *userRepositoryImpl) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	user := &domain.User{}
	query := "SELECT id, username, password, role FROM users WHERE username = $1"
	err := u.sqldb.GetContext(ctx, user, query, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser implements UserRepository.
func (u *userRepositoryImpl) CreateUser(ctx context.Context, user *domain.User) error {
	query := "INSERT INTO users (username, password, role) VALUES ($1, $2, $3)"
	_, err := u.sqldb.ExecContext(ctx, query, user.Username, user.Password, user.Role)
	return err
}

// DeleteUser implements UserRepository.
func (u *userRepositoryImpl) DeleteUser(ctx context.Context, id string) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := u.sqldb.ExecContext(ctx, query, id)
	return err
}

// GetUser implements UserRepository.
func (u *userRepositoryImpl) GetUser(ctx context.Context, id string) (*domain.User, error) {
	user := &domain.User{}
	query := "SELECT id, username, password, role FROM users WHERE id = $1"
	err := u.sqldb.GetContext(ctx, user, query, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser implements UserRepository.
func (u *userRepositoryImpl) UpdateUser(ctx context.Context, user *domain.User) error {
	tx, err := u.sqldb.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	query := "UPDATE users SET username = $1, role = $2, role WHERE id = $3"
	_, err = tx.ExecContext(ctx, query, user.Username, user.Role, user.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func NewUserRepository(sqlDb *sqlx.DB) UserRepository {
	return &userRepositoryImpl{sqldb: sqlDb}
}
