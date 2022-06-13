package usecase

import (
	"context"

	"github.com/quazar2000/my-go-clean-arch-temp/models"
)

type commentvoteUsecase struct {
	commentvoteRepository models.CommentVoteRepository
}

//  NewCommentVoteUsecase will create new comment vote Usecase object representation of models.CommentVoteUsecase interface
func NewCommentVoteUsecase(cv models.CommentVoteRepository) models.CommentVoteUsecase {
	return &commentvoteUsecase{
		commentvoteRepository: cv,
	}
}

func (cv *commentvoteUsecase) Like(ctx context.Context, commentVote *models.CommentVote) (err error) {
	return nil
}

func (cv *commentvoteUsecase) Dislike(ctx context.Context, commentVote *models.CommentVote) (err error) {
	return nil
}

func (cv *commentvoteUsecase) Delete(ctx context.Context, id int) (err error) {
	return nil
}
