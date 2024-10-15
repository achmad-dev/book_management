package repository

import (
	"github.com/achmad-dev/internal/category/internal/domain"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository interface {
	CreateCategory(name string) error
	GetCategory(id string) (*domain.Category, error)
	GetCategoryByName(name string) (*domain.Category, error)
	UpdateCategory(id string, name string) error
	DeleteCategory(id string) error
	ListCategories() ([]*domain.Category, error)
}

type categoryRepositoryImpl struct {
	sqlDB *sqlx.DB
}

// CreateCategory implements CategoryRepository.
func (c *categoryRepositoryImpl) CreateCategory(name string) error {
	_, err := c.sqlDB.Exec("INSERT INTO category (name) VALUES ($1)", name)
	return err
}

// DeleteCategory implements CategoryRepository.
func (c *categoryRepositoryImpl) DeleteCategory(id string) error {
	_, err := c.sqlDB.Exec("DELETE FROM category WHERE id = $1", id)
	return err
}

// GetCategory implements CategoryRepository.
func (c *categoryRepositoryImpl) GetCategory(id string) (*domain.Category, error) {
	var category domain.Category
	err := c.sqlDB.Get(&category, "SELECT id, name FROM category WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// GetCategoryByName implements CategoryRepository.
func (c *categoryRepositoryImpl) GetCategoryByName(name string) (*domain.Category, error) {
	var category domain.Category
	err := c.sqlDB.Get(&category, "SELECT id, name FROM category WHERE name = $1", name)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// ListCategories implements CategoryRepository.
func (c *categoryRepositoryImpl) ListCategories() ([]*domain.Category, error) {
	var categories []*domain.Category
	err := c.sqlDB.Select(&categories, "SELECT id, name FROM category")
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// UpdateCategory implements CategoryRepository.
func (c *categoryRepositoryImpl) UpdateCategory(id string, name string) error {
	_, err := c.sqlDB.Exec("UPDATE category SET name = $1 WHERE id = $2", name, id)
	return err
}

func NewCategoryRepository(sqlDb *sqlx.DB) CategoryRepository {
	return &categoryRepositoryImpl{
		sqlDB: sqlDb,
	}
}
