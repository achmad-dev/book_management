package repository

import (
	"github.com/achmad-dev/internal/user/internal/domain"
	"github.com/jmoiron/sqlx"
)

type UserBorrowedBookRepository interface {
	CreateUserBorrowedBook(userID, bookID, title string, quantity int) error
	DeleteUserBorrowedBook(id string) error
	ReturnUserBorrowedBook(id string, quantity int) error
	GetUserBorrowedBook(id string) (*domain.UserBorrowedBook, error)
	GetUserBorrowedBookByUserID(userID string) ([]*domain.UserBorrowedBook, error)
}

type userBorrowedBookRepositoryImpl struct {
	sqlDB *sqlx.DB
}

// CreateUserBorrowedBook implements UserBorrowedBookRepository.
// Create or Update user borrowed book.
func (u *userBorrowedBookRepositoryImpl) CreateUserBorrowedBook(userID string, bookID string, title string, quantity int) error {
	query := `
		INSERT INTO user_borrowed_books (id, user_id, book_title, quantity)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (id) DO UPDATE
		SET quantity = user_borrowed_books.quantity + EXCLUDED.quantity,
			updated_at = now()
	`
	_, err := u.sqlDB.Exec(query, bookID, userID, title, quantity)
	return err
}

// DeleteUserBorrowedBook implements UserBorrowedBookRepository.
func (u *userBorrowedBookRepositoryImpl) DeleteUserBorrowedBook(id string) error {
	query := `DELETE FROM user_borrowed_books WHERE id = $1`
	_, err := u.sqlDB.Exec(query, id)
	return err
}

// GetUserBorrowedBook implements UserBorrowedBookRepository.
func (u *userBorrowedBookRepositoryImpl) GetUserBorrowedBook(id string) (*domain.UserBorrowedBook, error) {
	query := `SELECT id, user_id, book_title, quantity, created_at, updated_at FROM user_borrowed_books WHERE id = $1`
	var book domain.UserBorrowedBook
	err := u.sqlDB.Get(&book, query, id)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// GetUserBorrowedBookByUserID implements UserBorrowedBookRepository.
func (u *userBorrowedBookRepositoryImpl) GetUserBorrowedBookByUserID(userID string) ([]*domain.UserBorrowedBook, error) {
	query := `SELECT id, user_id, book_title, quantity, created_at, updated_at FROM user_borrowed_books WHERE user_id = $1`
	var books []*domain.UserBorrowedBook
	err := u.sqlDB.Select(&books, query, userID)
	if err != nil {
		return nil, err
	}
	return books, nil
}

// ReturnUserBorrowedBook implements UserBorrowedBookRepository.
// Return user borrowed book.
// When quantity of borrowed book reach 0, delete the record.
func (u *userBorrowedBookRepositoryImpl) ReturnUserBorrowedBook(id string, quantity int) error {
	tx, err := u.sqlDB.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var currentQuantity int
	query := `SELECT quantity FROM user_borrowed_books WHERE id = $1`
	err = tx.Get(&currentQuantity, query, id)
	if err != nil {
		return err
	}

	newQuantity := currentQuantity - quantity
	if newQuantity <= 0 {
		query = `DELETE FROM user_borrowed_books WHERE id = $1`
		_, err = tx.Exec(query, id)
	} else {
		query = `UPDATE user_borrowed_books SET quantity = $1, updated_at = now() WHERE id = $2`
		_, err = tx.Exec(query, newQuantity, id)
	}

	return err
}

func NewUserBorrowedBookRepository(sqlDB *sqlx.DB) UserBorrowedBookRepository {
	return &userBorrowedBookRepositoryImpl{sqlDB: sqlDB}
}
