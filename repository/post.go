package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/quazar2000/my-go-clean-arch-temp/models"
)

type postgresPostRepository struct {
	Conn *pgxpool.Pool
}

// NewPostgresPostRepository will create an object that represent the post. Repository interface
func NewPostgresPostRepository(Conn *pgxpool.Pool) models.PostRepository {
	return &postgresPostRepository{
		Conn: Conn,
	}
}

// Create post ...
func (p *postgresPostRepository) Create(ctx context.Context, post *models.Post) (id int, err error) {
	return 0, nil
}

// Update post ...
func (p *postgresPostRepository) Update(ctx context.Context, post *models.Post) (err error) {
	return nil
}

// GetAll posts ...
func (p *postgresPostRepository) GetAll(ctx context.Context, post *models.Post) (posts *[]models.PostRequestDTO, err error) {
	return nil, nil
}

// GetbyID the post ...
func (p *postgresPostRepository) GetByID(ctx context.Context, id int) (post *models.PostRequestDTO, err error) {
	return nil, nil
}

// Delete post ...
func (p *postgresPostRepository) Delete(ctx context.Context, id int) (err error) {
	return nil
}
