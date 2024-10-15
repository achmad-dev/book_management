package service

import (
	"github.com/achmad-dev/internal/category/internal/domain"
	"github.com/achmad-dev/internal/category/internal/repository"
	"github.com/sirupsen/logrus"
)

type CategoryService interface {
	CreateCategory(name string) error
	GetCategory(id string) (*domain.Category, error)
	GetCategoryByName(name string) (*domain.Category, error)
	UpdateCategory(id string, name string) error
	DeleteCategory(id string) error
	ListCategories() ([]*domain.Category, error)
}

type categoryServiceImpl struct {
	ctRepo repository.CategoryRepository
	log    *logrus.Logger
}

// CreateCategory implements CategoryService.
func (c *categoryServiceImpl) CreateCategory(name string) error {
	err := c.ctRepo.CreateCategory(name)
	if err != nil {
		c.log.Errorf("Failed to create category: %v", err)
		return err
	}
	c.log.Infof("Category created successfully: %s", name)
	return nil
}

// DeleteCategory implements CategoryService.
func (c *categoryServiceImpl) DeleteCategory(id string) error {
	err := c.ctRepo.DeleteCategory(id)
	if err != nil {
		c.log.Errorf("Failed to delete category with id %s: %v", id, err)
		return err
	}
	c.log.Infof("Category deleted successfully: %s", id)
	return nil
}

// GetCategory implements CategoryService.
func (c *categoryServiceImpl) GetCategory(id string) (*domain.Category, error) {
	category, err := c.ctRepo.GetCategory(id)
	if err != nil {
		c.log.Errorf("Failed to get category with id %s: %v", id, err)
		return nil, err
	}
	c.log.Infof("Category retrieved successfully: %+v", category)
	return category, nil
}

// GetCategoryByName implements CategoryService.
func (c *categoryServiceImpl) GetCategoryByName(name string) (*domain.Category, error) {
	category, err := c.ctRepo.GetCategoryByName(name)
	if err != nil {
		c.log.Errorf("Failed to get category with name %s: %v", name, err)
		return nil, err
	}
	c.log.Infof("Category retrieved successfully: %+v", category)
	return category, nil
}

// ListCategories implements CategoryService.
func (c *categoryServiceImpl) ListCategories() ([]*domain.Category, error) {
	categories, err := c.ctRepo.ListCategories()
	if err != nil {
		c.log.Errorf("Failed to list categories: %v", err)
		return nil, err
	}
	c.log.Infof("Categories listed successfully: %+v", categories)
	return categories, nil
}

// UpdateCategory implements CategoryService.
func (c *categoryServiceImpl) UpdateCategory(id string, name string) error {
	err := c.ctRepo.UpdateCategory(id, name)
	if err != nil {
		c.log.Errorf("Failed to update category with id %s: %v", id, err)
		return err
	}
	c.log.Infof("Category updated successfully: %s", id)
	return nil
}

func NewCategoryService(ctRepo repository.CategoryRepository, log *logrus.Logger) CategoryService {
	return &categoryServiceImpl{ctRepo: ctRepo, log: log}
}
