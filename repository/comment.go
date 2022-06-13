package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/quazar2000/my-go-clean-arch-temp/models"
)

type postgresCommentRepository struct {
	Conn *pgxpool.Pool
}

// NewPostgresCommentRepository will create an object tat represent the post. Repository interface
func NewPostgresCommentRepository(Conn *pgxpool.Pool) models.CommentRepository {
	return &postgresCommentRepository{
		Conn: Conn,
	}
}

// Create comment ...
func (c *postgresCommentRepository) Create(ctx context.Context, comment *models.Comment) (id int, err error) {
	return 0, nil
}

// Update comment ...
func (c *postgresCommentRepository) Update(ctx context.Context, comment *models.Comment) (err error) {
	return nil
}

// GetByUserID comments ...
func (c *postgresCommentRepository) GetByUserID(ctx context.Context, user_id int) (comments []*models.CommentRequestDTO, err error) {
	return nil, nil
}

// GetByPostID comments ...
func (c *postgresCommentRepository) GetByPostID(ctx context.Context, post_id int) (comments []*models.CommentRequestDTO, err error) {
	return nil, nil
}

// Delete comment
func (c *postgresCommentRepository) Delete(ctx context.Context, id int) (err error) {
	return nil
}
