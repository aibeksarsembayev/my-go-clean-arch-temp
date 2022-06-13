package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/quazar2000/my-go-clean-arch-temp/models"
)

type postgresCategoryRepository struct {
	Conn *pgxpool.Pool
}

// NewPostgresCategoryRepository will create an onject that reposerent the category. Repository interface
func NewPostgresCategoryRepository(Conn *pgxpool.Pool) models.CategoryRepository {
	return &postgresCategoryRepository{
		Conn: Conn,
	}
}

// Create category ...
func (c *postgresCategoryRepository) Create(ctx context.Context, category *models.Category) (id int, err error) {
	return 0, nil
}

// GetAll categories ...
func (c *postgresCategoryRepository) GetAll(ctx context.Context, category *models.Category) (categories []*models.CategoryRequestDTO, err error) {
	return nil, nil
}

// GetbyID category ...
func (c *postgresCategoryRepository) GetbyID(ctx context.Context, id int) (category *models.CategoryRequestDTO, err error) {
	return nil, nil
}

// Update category ...
func (c *postgresCategoryRepository) Update(ctx context.Context, category *models.Category) (err error) {
	return nil
}

// Delete category ...
func (c *postgresCategoryRepository) Delete(ctx context.Context, id int) (err error) {
	return nil
}
