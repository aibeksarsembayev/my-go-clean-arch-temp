package usecase

import (
	"context"

	"github.com/quazar2000/my-go-clean-arch-temp/models"
)

type categoryUsecase struct {
	categoryRepository models.CategoryRepository
}

// NewCategoryUsecase will create new an category Usecase object representation of models.UserUsecase interface
func NewCategoryUsecase(c models.CategoryRepository) models.CategoryUsecase {
	return &categoryUsecase{
		categoryRepository: c,
	}
}

// Create category ...
func (c *categoryUsecase) Create(ctx context.Context, category *models.Category) (id int, err error) {
	return 0, nil
}

// GetAll categories ...
func (c *categoryUsecase) GetAll(ctx context.Context, category *models.Category) (categories []*models.CategoryRequestDTO, err error) {
	return nil, nil
}

// GetbyID category ...
func (c *categoryUsecase) GetbyID(ctx context.Context, id int) (category *models.CategoryRequestDTO, err error) {
	return nil, nil
}

// Update category ...
func (c *categoryUsecase) Update(ctx context.Context, category *models.Category) (err error) {
	return nil
}

// Delete category ...
func (c *categoryUsecase) Delete(ctx context.Context, id int) (err error) {
	return nil
}
