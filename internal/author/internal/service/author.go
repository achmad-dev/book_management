package service

import (
	"github.com/achmad-dev/internal/author/internal/domain"
	"github.com/achmad-dev/internal/author/internal/repository"
	"github.com/sirupsen/logrus"
)

type AuthorService interface {
	GetAuthor(id string) (*domain.Author, error)
	GetAuthorByName(name string) (*domain.Author, error)
	CreateAuthor(author *domain.Author) error
	UpdateAuthor(id string, author *domain.Author) error
	DeleteAuthor(id string) error
	DeleteAuthorByName(name string) error
	ListAuthors() ([]*domain.Author, error)
}

type authorServiceImpl struct {
	authorRepo repository.AuthorRepository
	log        *logrus.Logger
}

// CreateAuthor implements AuthorService.
func (a *authorServiceImpl) CreateAuthor(author *domain.Author) error {
	a.log.Infof("Creating author: %v", author)
	err := a.authorRepo.CreateAuthor(author)
	if err != nil {
		a.log.Errorf("Error creating author: %v", err)
	}
	return err
}

// DeleteAuthor implements AuthorService.
func (a *authorServiceImpl) DeleteAuthor(id string) error {
	a.log.Infof("Deleting author with ID: %s", id)
	err := a.authorRepo.DeleteAuthor(id)
	if err != nil {
		a.log.Errorf("Error deleting author with ID %s: %v", id, err)
	}
	return err
}

// GetAuthor implements AuthorService.
func (a *authorServiceImpl) GetAuthor(id string) (*domain.Author, error) {
	a.log.Infof("Getting author with ID: %s", id)
	author, err := a.authorRepo.GetAuthor(id)
	if err != nil {
		a.log.Errorf("Error getting author with ID %s: %v", id, err)
	}
	return author, err
}

// GetAuthorByName implements AuthorService.
func (a *authorServiceImpl) GetAuthorByName(name string) (*domain.Author, error) {
	a.log.Infof("Getting author with name: %s", name)
	author, err := a.authorRepo.GetAuthorByName(name)
	if err != nil {
		a.log.Errorf("Error getting author with name %s: %v", name, err)
	}
	return author, err
}

// UpdateAuthor implements AuthorService.
func (a *authorServiceImpl) UpdateAuthor(id string, author *domain.Author) error {
	a.log.Infof("Updating author with ID: %s", id)
	err := a.authorRepo.UpdateAuthor(id, author)
	if err != nil {
		a.log.Errorf("Error updating author with ID %s: %v", id, err)
	}
	return err
}

// DeleteAuthorByName implements AuthorService.
func (a *authorServiceImpl) DeleteAuthorByName(name string) error {
	a.log.Infof("Deleting author with name: %s", name)
	err := a.authorRepo.DeleteAuthorByName(name)
	if err != nil {
		a.log.Errorf("Error deleting author with name %s: %v", name, err)
	}
	return err
}

// ListAuthors implements AuthorService.
func (a *authorServiceImpl) ListAuthors() ([]*domain.Author, error) {
	a.log.Info("Listing all authors")
	authors, err := a.authorRepo.ListAuthors()
	if err != nil {
		a.log.Errorf("Error listing authors: %v", err)
	}
	return authors, err
}

func NewAuthorService(authorRepo repository.AuthorRepository, log *logrus.Logger) AuthorService {
	return &authorServiceImpl{
		authorRepo: authorRepo,
		log:        log,
	}
}
