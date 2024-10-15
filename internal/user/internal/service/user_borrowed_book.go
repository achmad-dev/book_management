package service

import (
	"github.com/achmad-dev/internal/user/internal/domain"
	"github.com/achmad-dev/internal/user/internal/repository"
	"github.com/sirupsen/logrus"
)

type UserBorrowedBookService interface {
	BorrowBook(userID, bookID, title string, quantity int) error
	ReturnBook(id string, quantity int) error
	GetBorrowedBook(id string) (*domain.UserBorrowedBook, error)
	GetBorrowedBooksByUserID(userID string) ([]*domain.UserBorrowedBook, error)
}

type userBorrowedBookServiceImpl struct {
	repo repository.UserBorrowedBookRepository
	log  *logrus.Logger
}

func (s *userBorrowedBookServiceImpl) BorrowBook(userID, bookID, title string, quantity int) error {
	s.log.Infof("Borrowing book: userID=%s, bookID=%s, title=%s, quantity=%d", userID, bookID, title, quantity)
	err := s.repo.CreateUserBorrowedBook(userID, bookID, title, quantity)
	if err != nil {
		s.log.Errorf("Error borrowing book: %v", err)
	}
	return err
}

func (s *userBorrowedBookServiceImpl) ReturnBook(id string, quantity int) error {
	s.log.Infof("Returning book: id=%s, quantity=%d", id, quantity)
	err := s.repo.ReturnUserBorrowedBook(id, quantity)
	if err != nil {
		s.log.Errorf("Error returning book: %v", err)
	}
	return err
}

func (s *userBorrowedBookServiceImpl) GetBorrowedBook(id string) (*domain.UserBorrowedBook, error) {
	s.log.Infof("Getting borrowed book: id=%s", id)
	book, err := s.repo.GetUserBorrowedBook(id)
	if err != nil {
		s.log.Errorf("Error getting borrowed book: %v", err)
	}
	return book, err
}

func (s *userBorrowedBookServiceImpl) GetBorrowedBooksByUserID(userID string) ([]*domain.UserBorrowedBook, error) {
	s.log.Infof("Getting borrowed books by userID: %s", userID)
	books, err := s.repo.GetUserBorrowedBookByUserID(userID)
	if err != nil {
		s.log.Errorf("Error getting borrowed books by userID: %v", err)
	}
	return books, err
}

func NewUserBorrowedBookService(repo repository.UserBorrowedBookRepository, log *logrus.Logger) UserBorrowedBookService {
	return &userBorrowedBookServiceImpl{repo: repo, log: log}
}
