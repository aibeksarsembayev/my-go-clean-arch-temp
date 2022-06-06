package usecase

import (
	"context"

	"github.com/quazar2000/my-go-clean-arch-temp/models"
)

type postUsecase struct {
	postRepository models.PostRepository
}

// NewPostUsecase will create new an post Usecase object representation of models.PostUsecase interface
func NewPostUsecase(p models.PostRepository) models.PostUsecase {
	return &postUsecase{
		postRepository: p,
	}
}

// Create post ...
func (p *postUsecase) Create(ctx context.Context, post *models.Post) (id int, err error) {
	return 0, nil
}

// Update post ...
func (p *postUsecase) Update(ctx context.Context, post *models.Post) (err error) {
	return nil
}

// GetAll posts ...
func (p *postUsecase) GetAll(ctx context.Context, post *models.Post) (posts *[]models.PostRequestDTO, err error) {
	return nil, nil
}

// GetbyID the post ...
func (p *postUsecase) GetByID(ctx context.Context, id int) (post *models.PostRequestDTO, err error) {
	return nil, nil
}

// Delete post ...
func (p *postUsecase) Delete(ctx context.Context, id int) (err error) {
	return nil
}
