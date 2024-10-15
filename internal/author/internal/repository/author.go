package repository

import (
	"time"

	"github.com/achmad-dev/internal/author/internal/domain"
	"github.com/jmoiron/sqlx"
)

type AuthorRepository interface {
	GetAuthor(id string) (*domain.Author, error)
	GetAuthorByName(name string) (*domain.Author, error)
	CreateAuthor(author *domain.Author) error
	UpdateAuthor(id string, author *domain.Author) error
	DeleteAuthor(id string) error
	DeleteAuthorByName(name string) error
	ListAuthors() ([]*domain.Author, error)
}

type authorRepositoryImpl struct {
	sqlDb *sqlx.DB
}

// CreateAuthor implements AuthorRepository.
func (a *authorRepositoryImpl) CreateAuthor(author *domain.Author) error {
	query := `INSERT INTO authors (name) VALUES ($1)`
	_, err := a.sqlDb.Exec(query, author.Name)
	return err
}

// DeleteAuthor implements AuthorRepository.
func (a *authorRepositoryImpl) DeleteAuthor(id string) error {
	query := `DELETE FROM authors WHERE id = $1`
	_, err := a.sqlDb.Exec(query, id)
	return err
}

// DeleteAuthorByName implements AuthorRepository.
func (a *authorRepositoryImpl) DeleteAuthorByName(name string) error {
	query := `DELETE FROM authors WHERE name = $1`
	_, err := a.sqlDb.Exec(query, name)
	return err
}

// GetAuthor implements AuthorRepository.
func (a *authorRepositoryImpl) GetAuthor(id string) (*domain.Author, error) {
	query := `SELECT id, name, created_at, updated_at FROM authors WHERE id = $1`
	var author domain.Author
	err := a.sqlDb.Get(&author, query, id)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

// GetAuthorByName implements AuthorRepository.
func (a *authorRepositoryImpl) GetAuthorByName(name string) (*domain.Author, error) {
	query := `SELECT id, name, created_at, updated_at FROM authors WHERE name = $1`
	var author domain.Author
	err := a.sqlDb.Get(&author, query, name)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

// UpdateAuthor implements AuthorRepository.
func (a *authorRepositoryImpl) UpdateAuthor(id string, author *domain.Author) error {
	tx, err := a.sqlDb.Beginx()
	if err != nil {
		return err
	}

	query := `UPDATE authors SET name = $1, updated_at = $2 WHERE id = $3`
	_, err = tx.Exec(query, author.Name, time.Now().UTC(), id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// ListAuthors implements AuthorRepository.
func (a *authorRepositoryImpl) ListAuthors() ([]*domain.Author, error) {
	query := `SELECT id, name, created_at, updated_at FROM authors`
	var authors []*domain.Author
	err := a.sqlDb.Select(&authors, query)
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func NewAuthorRepository(sqlDb *sqlx.DB) AuthorRepository {
	return &authorRepositoryImpl{
		sqlDb: sqlDb,
	}
}
