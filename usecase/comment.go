package usecase

import (
	"context"

	"github.com/quazar2000/my-go-clean-arch-temp/models"
)

type commentUsecase struct {
	commentRepository models.CommentRepository
}

// NewCommentUsecase will create new comment Usecase object representation of models.CommentUsecase interface
func NewCommentUsecase(c models.CommentRepository) models.CommentUsecase {
	return &commentUsecase{
		commentRepository: c,
	}
}

// Create comment ...
func (c *commentUsecase) Create(ctx context.Context, comment *models.Comment) (id int, err error) {
	return 0, nil
}

// Update comment ...
func (c *commentUsecase) Update(ctx context.Context, comment *models.Comment) (err error) {
	return nil
}

// GetByUserID comments ...
func (c *commentUsecase) GetByUserID(ctx context.Context, user_id int) (comments []*models.CommentRequestDTO, err error) {
	return nil, nil
}

// GetByPostID comments ...
func (c *commentUsecase) GetByPostID(ctx context.Context, post_id int) (comments []*models.CommentRequestDTO, err error) {
	return nil, nil
}

// Delete comment
func (c *commentUsecase) Delete(ctx context.Context, id int) (err error) {
	return nil
}
