package repository

import (
	"errors"

	"github.com/achmad-dev/internal/book/internal/domain"
	"github.com/jmoiron/sqlx"
)

type BookRepository interface {
	GetBook(id string) (*domain.Book, error)
	GetBookByTitle(title string) (*domain.Book, error)
	GetBooksByAuthorName(author string) ([]*domain.Book, error)
	GetPopularBooksByCategory(category string) ([]*domain.Book, error)
	BorrowBook(id string, quantity int) error
	ReturnBook(id string, quantity int) error
	CreateBook(book *domain.Book) error
	UpdateBook(id string, book *domain.Book) error
	DeleteBook(id string) error
	ListBooks() ([]*domain.Book, error)
}

type bookRepositoryImpl struct {
	sqlDb *sqlx.DB
}

// implements BookRepository.
// When borrowed book quantity is more than available stock, it should return an error.
// When borrowed time is above 3 times, update the book is_popular to true.
func (b *bookRepositoryImpl) BorrowBook(id string, quantity int) error {
	tx, err := b.sqlDb.Beginx()
	if err != nil {
		return err
	}

	var stock, borrowed int
	err = tx.Get(&stock, "SELECT stock FROM books WHERE id=$1", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if stock < quantity {
		tx.Rollback()
		return errors.New("not enough stock")
	}

	_, err = tx.Exec("UPDATE books SET stock = stock - $1, borrowed = borrowed + 1 WHERE id = $2", quantity, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Get(&borrowed, "SELECT borrowed FROM books WHERE id=$1", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if borrowed > 3 {
		_, err = tx.Exec("UPDATE books SET is_popular = TRUE WHERE id = $1", id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

// ReturnBook implements BookRepository.
func (b *bookRepositoryImpl) ReturnBook(id string, quantity int) error {
	tx, err := b.sqlDb.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE books SET stock = stock + $1 WHERE id = $2", quantity, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// CreateBook implements BookRepository.
func (b *bookRepositoryImpl) CreateBook(book *domain.Book) error {
	query := `
		INSERT INTO books (
			author_id, category_id, title, author, category, stock
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
	`
	_, err := b.sqlDb.Exec(query, book.Author_id, book.Category_id, book.Title, book.Author, book.Category, book.Stock)
	return err
}

// DeleteBook implements BookRepository.
func (b *bookRepositoryImpl) DeleteBook(id string) error {
	tx, err := b.sqlDb.Beginx()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// GetBook implements BookRepository.
func (b *bookRepositoryImpl) GetBook(id string) (*domain.Book, error) {
	var book domain.Book
	err := b.sqlDb.Get(&book, "SELECT * FROM books WHERE id = $1", id)
	return &book, err
}

// GetBookByTitle implements BookRepository.
func (b *bookRepositoryImpl) GetBookByTitle(title string) (*domain.Book, error) {
	var book domain.Book
	err := b.sqlDb.Get(&book, "SELECT * FROM books WHERE title = $1", title)
	return &book, err
}

// GetBooksByAuthorName implements BookRepository.
func (b *bookRepositoryImpl) GetBooksByAuthorName(author string) ([]*domain.Book, error) {
	var books []*domain.Book
	err := b.sqlDb.Select(&books, "SELECT * FROM books WHERE author = $1", author)
	return books, err
}

// GetPopularBooksByCategory implements BookRepository.
// Get popular books where is_popular is true and category is the same as the input.
func (b *bookRepositoryImpl) GetPopularBooksByCategory(category string) ([]*domain.Book, error) {
	var books []*domain.Book
	err := b.sqlDb.Select(&books, "SELECT * FROM books WHERE is_popular = TRUE AND category = $1", category)
	return books, err
}

// ListBooks implements BookRepository.
func (b *bookRepositoryImpl) ListBooks() ([]*domain.Book, error) {
	var books []*domain.Book
	err := b.sqlDb.Select(&books, "SELECT * FROM books")
	return books, err
}

// UpdateBook implements BookRepository.
func (b *bookRepositoryImpl) UpdateBook(id string, book *domain.Book) error {
	_, err := b.sqlDb.NamedExec(`UPDATE books SET title=:title, author=:author, category=:category, stock=:stock, borrowed=:borrowed, is_popular=:is_popular, updated_at=:updated_at WHERE id=:id`, book)
	return err
}

func NewBookRepository(sqlDb *sqlx.DB) BookRepository {
	return &bookRepositoryImpl{
		sqlDb: sqlDb,
	}
}
