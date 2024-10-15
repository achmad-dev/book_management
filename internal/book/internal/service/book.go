package service

import (
	"github.com/achmad-dev/internal/book/internal/domain"
	"github.com/achmad-dev/internal/book/internal/repository"
	"github.com/sirupsen/logrus"
)

type BookService interface {
	GetBook(id string) (*domain.Book, error)
	GetBookByTitle(title string) (*domain.Book, error)
	GetBooksByAuthorName(author string) ([]*domain.Book, error)
	GetPopularBooksByCategory(category string) ([]*domain.Book, error)
	CreateBook(book *domain.Book) error
	UpdateBook(id string, book *domain.Book) error
	DeleteBook(id string) error
	ListBooks() ([]*domain.Book, error)
	BorrowBook(id string, quantity int) error
	ReturnBook(id string, quantity int) error
}

type bookServiceImpl struct {
	bookRepo repository.BookRepository
	log      *logrus.Logger
}

// CreateBook implements BookService.
func (b *bookServiceImpl) CreateBook(book *domain.Book) error {
	b.log.Info("Creating book: ", book.Title)
	err := b.bookRepo.CreateBook(book)
	if err != nil {
		b.log.Error("Error creating book: ", err)
	}
	return err
}

// DeleteBook implements BookService.
func (b *bookServiceImpl) DeleteBook(id string) error {
	b.log.Info("Deleting book with ID: ", id)
	err := b.bookRepo.DeleteBook(id)
	if err != nil {
		b.log.Error("Error deleting book: ", err)
	}
	return err
}

// GetBook implements BookService.
func (b *bookServiceImpl) GetBook(id string) (*domain.Book, error) {
	b.log.Info("Fetching book with ID: ", id)
	book, err := b.bookRepo.GetBook(id)
	if err != nil {
		b.log.Error("Error fetching book: ", err)
	}
	return book, err
}

// GetBookByTitle implements BookService.
func (b *bookServiceImpl) GetBookByTitle(title string) (*domain.Book, error) {
	b.log.Info("Fetching book with title: ", title)
	book, err := b.bookRepo.GetBookByTitle(title)
	if err != nil {
		b.log.Error("Error fetching book: ", err)
	}
	return book, err
}

// GetBooksByAuthorName implements BookService.
func (b *bookServiceImpl) GetBooksByAuthorName(author string) ([]*domain.Book, error) {
	b.log.Info("Fetching books by author: ", author)
	books, err := b.bookRepo.GetBooksByAuthorName(author)
	if err != nil {
		b.log.Error("Error fetching books: ", err)
	}
	return books, err
}

// GetPopularBooksByCategory implements BookService.
func (b *bookServiceImpl) GetPopularBooksByCategory(category string) ([]*domain.Book, error) {
	b.log.Info("Fetching popular books in category: ", category)
	books, err := b.bookRepo.GetPopularBooksByCategory(category)
	if err != nil {
		b.log.Error("Error fetching popular books: ", err)
	}
	return books, err
}

// ListBooks implements BookService.
func (b *bookServiceImpl) ListBooks() ([]*domain.Book, error) {
	b.log.Info("Listing all books")
	books, err := b.bookRepo.ListBooks()
	if err != nil {
		b.log.Error("Error listing books: ", err)
	}
	return books, err
}

// UpdateBook implements BookService.
func (b *bookServiceImpl) UpdateBook(id string, book *domain.Book) error {
	b.log.Info("Updating book with ID: ", id)
	err := b.bookRepo.UpdateBook(id, book)
	if err != nil {
		b.log.Error("Error updating book: ", err)
	}
	return err
}

// BorrowBook implements BookService.
func (b *bookServiceImpl) BorrowBook(id string, quantity int) error {
	b.log.Info("Borrowing book with ID: ", id, " Quantity: ", quantity)
	err := b.bookRepo.BorrowBook(id, quantity)
	if err != nil {
		b.log.Error("Error borrowing book: ", err)
	}
	return err
}

// ReturnBook implements BookService.
func (b *bookServiceImpl) ReturnBook(id string, quantity int) error {
	b.log.Info("Returning book with ID: ", id, " Quantity: ", quantity)
	err := b.bookRepo.ReturnBook(id, quantity)
	if err != nil {
		b.log.Error("Error returning book: ", err)
	}
	return err
}

func NewBookService(bookRepo repository.BookRepository, log *logrus.Logger) BookService {
	return &bookServiceImpl{bookRepo: bookRepo, log: log}
}
