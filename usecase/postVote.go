package usecase

import (
	"context"

	"github.com/quazar2000/my-go-clean-arch-temp/models"
)

type postvoteUsecase struct {
	postvoteRepository models.PostVoteRepository
}

// NewPostVoteUsecase will create new postvote Usecase object representation of models.PostVoteUsecase interface
func NewPostVoteUsecase(pv models.PostVoteRepository) models.PostVoteUsecase {
	return &postvoteUsecase{
		postvoteRepository: pv,
	}
}

// Like post vote ...
func (pv *postvoteUsecase) Like(ctx context.Context, postVote *models.PostVote) (err error) {
	return nil
}

// Dislike post vote ...
func (pv *postvoteUsecase) Dislike(ctx context.Context, postVote *models.PostVote) (err error) {
	return nil
}

// Delete post vote ...
func (pv *postvoteUsecase) Delete(ctx context.Context, id int) (err error) {
	return nil
}
